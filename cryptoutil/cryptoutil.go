package cryptoutil

import cryptorand "crypto/rand"

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
