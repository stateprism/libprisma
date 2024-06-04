package encryption

import (
	aes "crypto/aes"
	"crypto/cipher"
	"crypto/sha512"
	"github.com/stateprism/libprisma/cryptoutil"
	"github.com/stateprism/libprisma/cryptoutil/kdf"
)

type SecureAES struct {
	iv    []byte
	nonce []byte
	key   []byte
	salt  []byte
	iAes  cipher.Block
	enc   cipher.BlockMode
	dec   cipher.BlockMode
}

func NewSecureAES(key []byte) (*SecureAES, error) {
	derived := kdf.PbKdf2.Key(key, 4096, 32, sha512.New)
	key = derived.GetKey()
	salt := derived.GetSalt()
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
		salt: salt,
		iAes: bc,
		enc:  enc,
		dec:  dec,
	}, nil
}

func (s *SecureAES) EncryptInPlace(data []byte) error {
	// Ensure that the data is a multiple of the block size
	if len(data)%aes.BlockSize != 0 {
		data, _ = cryptoutil.Pad(data, aes.BlockSize)
	}
	blocks := len(data) / aes.BlockSize
	for i := range blocks {
		s.enc.CryptBlocks(data[i*aes.BlockSize:], data[i*aes.BlockSize:])
	}
	return nil
}

func (s *SecureAES) DecryptInPlace(data []byte, plain []byte) error {
	blocks := len(data) / aes.BlockSize
	for i := range blocks {
		s.dec.CryptBlocks(plain[i*aes.BlockSize:], data[i*aes.BlockSize:])
	}
	data, err := cryptoutil.Unpad(data, aes.BlockSize)
	if err != nil {
		return err
	}
	return nil
}

func (s *SecureAES) Encrypt(data []byte) ([]byte, error) {
	// Ensure that the data is a multiple of the block size
	if len(data)%aes.BlockSize != 0 {
		data, _ = cryptoutil.Pad(data, aes.BlockSize)
	}
	encrypted := make([]byte, len(data))
	s.enc.CryptBlocks(encrypted, data)
	return encrypted, nil
}

func (s *SecureAES) Decrypt(data []byte) ([]byte, error) {
	decrypted := make([]byte, len(data))
	s.dec.CryptBlocks(decrypted, data)
	decrypted, err := cryptoutil.Unpad(decrypted, aes.BlockSize)
	if err != nil {
		return nil, err
	}
	return decrypted, nil
}
