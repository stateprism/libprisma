package encryption_test

import (
	"encoding/base64"
	"github.com/google/go-cmp/cmp"
	"github.com/stateprism/libprisma/cryptoutil/encryption"
	"testing"
)

func TestSecureAES(t *testing.T) {
	randomData := []byte("This is a test string, it is not very long but it is long enough to test the encryption and decryption functions")
	secureAes, err := encryption.NewSecureAES([]byte("superSecretKey"))
	if err != nil {
		t.Errorf("Error creating SecureAES: %v", err)
	}

	encrypted, err := secureAes.EncryptToBytes(randomData)
	if err != nil {
		t.Errorf("Error encrypting data: %v", err)
	}

	secureAes, err = encryption.NewSecureAES([]byte("superSecretKey"))
	if err != nil {
		t.Errorf("Error creating SecureAES: %v", err)
	}

	decrypted, err := secureAes.DecryptFromBytes(encrypted)
	if err != nil {
		t.Errorf("Error decrypting data: %v", err)
	}

	encStr := base64.StdEncoding.EncodeToString(encrypted)

	secureAes, err = encryption.NewSecureAES([]byte("superSecretKey"))
	if err != nil {
		t.Errorf("Error creating SecureAES: %v", err)
	}

	encryptedData, _ := base64.StdEncoding.DecodeString(encStr)

	if !cmp.Equal(encrypted, encryptedData) == false {
		t.Errorf("Data wrong after base64 encoding/decoding")
	}

	decrypted, err = secureAes.DecryptFromBytes(encryptedData)
	if err != nil {
		t.Errorf("Error decrypting data: %v", err)
	}

	if cmp.Equal(randomData, decrypted) == false {
		t.Errorf("Decrypted data does not match original data")
	}

	// test detect tampering
	encrypted[0] = encrypted[0] ^ 0xFF
	_, err = secureAes.DecryptFromBytes(encrypted)
	if err == nil {
		t.Errorf("Tampered data was decrypted successfully")
	}
}
