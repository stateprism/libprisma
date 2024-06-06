package main

import (
	"encoding/base64"
	"fmt"
	"github.com/stateprism/libprisma/cryptoutil"
	"github.com/stateprism/libprisma/cryptoutil/encryption"
)

func main() {
	// aes shenanigans

	data := []byte("hello world, this is a test message!")
	key := cryptoutil.NewRandom(32)
	secureAES, err := encryption.NewSecureAES(key)
	if err != nil {
		panic(err)
	}
	encrypted, err := secureAES.Encrypt(data)
	tag := secureAES.Finish()
	if err != nil {
		panic(err)
	}
	fmt.Println("Encrypted:", base64.StdEncoding.EncodeToString(encrypted))

	decrypted, err := secureAES.Decrypt(encrypted, tag)
	if err != nil {
		panic(err)
	}

	fmt.Println("Decrypted:", string(decrypted))
}
