package cryptoutil

import (
	cryptorand "crypto/rand"
	"errors"
	"fmt"
)

// NewRandom generates a random salt of the specified length and returns it as a Salt type. If an error occurs
// while generating the salt or the generated salt length does not match the specified length, a panic is raised
// with an error message.
func NewRandom(l int) []byte {
	salt := make([]byte, l)
	n, err := cryptorand.Read(salt)
	if n != l || err != nil {
		panic("Error getting randomness, check your OS true randomness source!")
	}
	return salt
}

// Pad pads the data to the specified block size, if data is already the right size, this is a no-op
func Pad(data []byte, blockSize int) ([]byte, error) {
	if data == nil || len(data) == 0 {
		return nil, errors.New("empty data cannot be padded")
	}
	if len(data) > blockSize {
		return nil, errors.New("block size cannot be less than the data passed")
	}
	if len(data)%blockSize == 0 {
		return data, nil
	}
	neededPadding := blockSize - len(data)%blockSize
	if neededPadding > 255 {
		return nil, errors.New("cannot pad more than 255 bytes")
	}
	padding := make([]byte, neededPadding)
	for i := 0; i < neededPadding; i++ {
		padding[i] = byte(neededPadding)
	}
	return append(data, padding...), nil
}

func Unpad(data []byte, blockSize int) ([]byte, error) {
	if data == nil || len(data) == 0 {
		return nil, errors.New("empty data cannot be unpadded")
	} else if len(data) < blockSize {
		return nil, fmt.Errorf("data is shorter than block size")
	} else if len(data)%blockSize != 0 {
		return data, fmt.Errorf("cannot unpad data that's not a multiple of %d", blockSize)
	}
	padLen := int(data[len(data)-1])
	// look at the last half of the data
	if padLen > blockSize {
		return nil, fmt.Errorf("padding length is greater than block size")
	} else if padLen == 0 {
		return nil, fmt.Errorf("padding length is zero")
	} else if padLen > len(data) {
		return nil, fmt.Errorf("padding length is greater than data length")
	}

	return data[:len(data)-padLen], nil
}
