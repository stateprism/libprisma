package encryption

import (
	aes "crypto/aes"
	"crypto/cipher"
	"crypto/sha512"
	"errors"
	"github.com/stateprism/libprisma/cryptoutil"
	"github.com/stateprism/libprisma/cryptoutil/kdf"
	"hash"
)

type EncryptionError int

const (
	ErrInvalidKeyLength EncryptionError = iota
)

func (e EncryptionError) Error() string {
	switch e {
	case ErrInvalidKeyLength:
		return "invalid key length"
	default:
		return "unknown error"
	}
}

type SecureAES struct {
	iv   []byte
	key  []byte
	iAes cipher.Block
	enc  cipher.BlockMode
	dec  cipher.BlockMode
	h    hash.Hash
}

func NewSecureAES(key []byte) (*SecureAES, error) {
	derived := kdf.PbKdf2.Key(key, 4096, 32, sha512.New)
	key = derived.GetKey()
	iv := cryptoutil.NewRandom(aes.BlockSize)
	bc, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	enc := cipher.NewCBCEncrypter(bc, iv)
	dec := cipher.NewCBCDecrypter(bc, iv)
	h := sha512.New()
	h.Write(key)

	return &SecureAES{
		iv:   iv,
		key:  key,
		iAes: bc,
		enc:  enc,
		dec:  dec,
		h:    h,
	}, nil
}

func NewSecureAESWithSafeKey(key []byte) (*SecureAES, error) {
	iv := cryptoutil.NewRandom(aes.BlockSize)
	if len(key) != 32 {
		return nil, ErrInvalidKeyLength
	}

	bc, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	enc := cipher.NewCBCEncrypter(bc, iv)
	dec := cipher.NewCBCDecrypter(bc, iv)
	h := sha512.New()

	return &SecureAES{
		iv:   iv,
		key:  key,
		iAes: bc,
		enc:  enc,
		dec:  dec,
		h:    h,
	}, nil
}

func (s *SecureAES) GetKey() []byte {
	return s.key
}

func (s *SecureAES) GetIV() []byte {
	return s.iv
}

func (s *SecureAES) SetIV(iv []byte) bool {
	if len(iv) != aes.BlockSize {
		return false
	}
	s.iv = iv
	s.enc = cipher.NewCBCEncrypter(s.iAes, iv)
	s.dec = cipher.NewCBCDecrypter(s.iAes, iv)
	return true
}

func (s *SecureAES) Encrypt(data []byte) ([]byte, error) {
	s.h.Reset()
	s.h.Write(s.key)
	out := make([]byte, len(data)+(aes.BlockSize-len(data)%aes.BlockSize))
	outBlocker := cryptoutil.NewBlocker(aes.BlockSize, out)
	blocker := cryptoutil.NewBlocker(aes.BlockSize, data)
	for {
		_, encrypted := outBlocker.Next()
		n, block := blocker.Next()
		if n == 0 {
			break
		}
		if n < aes.BlockSize {
			block, _ = cryptoutil.Pad(block, aes.BlockSize)
		}
		s.enc.CryptBlocks(encrypted, block)
		s.h.Write(block)
	}
	return out, nil
}

func (s *SecureAES) Finish() []byte {
	var tag []byte
	tag = s.h.Sum(tag)
	s.h.Reset()
	return tag
}

func (s *SecureAES) GetTagSize() int {
	return s.h.Size()
}

func (s *SecureAES) GetIvSize() int {
	return s.iAes.BlockSize()
}

func (s *SecureAES) TagPlusIVSize() int {
	return s.iAes.BlockSize() + s.h.Size()
}

func (s *SecureAES) Decrypt(data []byte, tag []byte) ([]byte, error) {
	if len(tag) != s.h.Size() {
		return nil, errors.New("tag size mismatch")
	}
	s.h.Reset()
	s.h.Write(s.key)
	decrypted := make([]byte, len(data))
	decryptedBlocker := cryptoutil.NewBlocker(aes.BlockSize, decrypted)
	blocker := cryptoutil.NewBlocker(aes.BlockSize, data)
	for {
		_, decrypted := decryptedBlocker.Next()
		n, block := blocker.Next()
		if n == 0 {
			break
		}
		s.dec.CryptBlocks(decrypted, block)
		s.h.Write(decrypted)
	}
	decrypted, err := cryptoutil.Unpad(decrypted, aes.BlockSize)
	if err != nil {
		return nil, err
	}
	if !cryptoutil.SecureCompare(tag, s.Finish()) {
		return nil, errors.New("tag mismatch")
	}
	return decrypted, nil
}
