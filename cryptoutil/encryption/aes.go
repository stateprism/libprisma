package encryption

import (
	aes "crypto/aes"
	"crypto/cipher"
	"crypto/sha512"
	"github.com/stateprism/libprisma/cryptoutil"
	"github.com/stateprism/libprisma/cryptoutil/kdf"
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

	return &SecureAES{
		iv:   iv,
		key:  key,
		iAes: bc,
		enc:  enc,
		dec:  dec,
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

	return &SecureAES{
		iv:   iv,
		key:  key,
		iAes: bc,
		enc:  enc,
		dec:  dec,
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
	// Ensure that the data is a multiple of the block size
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
	}
	return out, nil
}

func (s *SecureAES) Decrypt(data []byte) ([]byte, error) {
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
	}
	decrypted, err := cryptoutil.Unpad(decrypted, aes.BlockSize)
	if err != nil {
		return nil, err
	}
	return decrypted, nil
}
