package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/stateprism/libprisma/cryptoutil"
	"github.com/stateprism/libprisma/cryptoutil/encryption"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(1)
		}
	}()

	func() {
		key := flag.String("key", "", "key to use")
		keyHex := flag.String("key-hex", "", "key in hex")
		modeEncrypt := flag.Bool("encrypt", false, "encrypt mode")
		modeDecrypt := flag.Bool("decrypt", false, "decrypt mode")
		tagIv := flag.String("tag-iv", "", "tag and iv in hex")
		outputFormat := flag.String("output-format", "hex", "output format, hex, base64 or raw")
		flag.Parse()

		if *modeEncrypt && *modeDecrypt {
			panic("cannot encrypt and decrypt at the same time")
		} else if *key != "" && *keyHex != "" {
			panic("provide only key or key-hex, not both")
		} else if *key == "" && *keyHex == "" {
			panic("provide key or key-hex")
		}

		var keyBytes []byte
		if *key != "" {
			keyBytes = []byte(*key)
		} else {
			k, err := hex.DecodeString(*keyHex)
			if err != nil {
				panic("invalid key hex")
			}
			keyBytes = k
		}

		if *modeEncrypt {
			encrypt(keyBytes, *outputFormat)
		} else if *modeDecrypt {
			decrypt(keyBytes, *tagIv, *outputFormat)
		} else {
			panic("no mode selected")
		}
	}()
}

func encrypt(key []byte, form string) {
	// stdin reader
	reader := bufio.NewReader(os.Stdin)
	// stdout writer
	writer := bufio.NewWriter(os.Stdout)

	cypher, err := encryption.NewSecureAES(key, encryption.AES256)
	if err != nil {
		panic(err)
	}
	var fFunc func([]byte) []byte
	switch form {
	case "hex":
		fFunc = func(b []byte) []byte {
			return []byte(hex.EncodeToString(b))
		}
	case "base64":
		fFunc = func(b []byte) []byte {
			return []byte(hex.EncodeToString(b))
		}
	case "raw":
		fFunc = func(b []byte) []byte {
			return b
		}
	default:
		panic("invalid output format")
	}
	// read input until end
	for {
		buff := make([]byte, 16)
		n, err := reader.Read(buff)
		if err != nil {
			break
		}
		line := buff[:n]
		// encrypt line
		encrypted, err := cypher.Encrypt(line)
		// write to stdout
		_, err = writer.Write(fFunc(encrypted))
		if err != nil {
			panic(err)
		}
	}
	err = writer.Flush()
	if err != nil {
		panic(err)
	}

	ivTag := cypher.GetIV()
	ivTag = append(ivTag, cypher.GetTag()...)

	// print tag and iv to stderr
	_, _ = fmt.Fprintf(os.Stderr, "\n\nYour tag and IV: %s\nKeep them safe :)\n", hex.EncodeToString(ivTag))

}

func decrypt(key []byte, ivTag string, form string) {
	// stdin reader
	reader := bufio.NewReader(os.Stdin)
	// stdout writer
	writer := bufio.NewWriter(os.Stdout)

	ivTagBytes, err := hex.DecodeString(ivTag)
	if err != nil {
		panic("invalid tag and iv hex")
	}
	iv := ivTagBytes[:16]
	tag := ivTagBytes[16:]

	cypher, err := encryption.NewSecureAES(key, encryption.AES256)
	if err != nil {
		panic(err)
	}

	cypher.SetIV(iv)

	var fFunc func([]byte) []byte
	switch form {
	case "hex":
		fFunc = func(b []byte) []byte {
			return []byte(hex.EncodeToString(b))
		}
	case "base64":
		fFunc = func(b []byte) []byte {
			return []byte(hex.EncodeToString(b))
		}
	case "raw":
		fFunc = func(b []byte) []byte {
			return b
		}
	default:
		panic("invalid output format")
	}
	// read input until end
	for {
		buff := make([]byte, 16)
		n, err := reader.Read(buff)
		if err != nil {
			break
		}
		line := buff[:n]
		// encrypt line
		decrypted, err := cypher.Decrypt(line)
		// write to stdout
		_, err = writer.Write(fFunc(decrypted))
		if err != nil {
			panic(err)
		}
	}
	err = writer.Flush()
	if err != nil {
		panic(err)
	}

	if cryptoutil.SecureCompare(tag, cypher.GetTag()) {
		fmt.Println("Tag match")
	} else {
		panic("Tag mismatch")
	}
}
