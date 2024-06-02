package kdf_test

import (
	"crypto/sha512"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/stateprism/libprisma/crypto/kdf"
	"hash"
	"testing"
)

func TestPbkdf2Key_Equals(t *testing.T) {
	someKey := "SomeKey"
	someKey2 := "SomeKey2"

	someKeyKdf := kdf.PbKdf2.Key(someKey, 4096, 32, sha512.New)

	if !someKeyKdf.Equals(someKey) {
		t.Error("Failed to properly compare keys")
	} else if someKeyKdf.Equals(someKey2) {
		t.Error("Failed to properly compare keys")
	}
}

func TestParsePbkdf2KeyFromString(t *testing.T) {
	someKey := "SomeKey"
	someKeyKdf := kdf.PbKdf2.Key(someKey, 4096, 32, sha512.New)
	someKeyStr := fmt.Sprintf("%s", someKeyKdf)
	v, _ := kdf.PbKdf2.FromString(someKeyStr)
	if !cmp.Equal(someKeyKdf, v) {
		t.Error("Decoded Key object from string doesn't match the original")
	}
}

func TestPbKdf2KeyBytes(t *testing.T) {
	tests := []struct {
		name string
		key  string
		iter int
		len  int
		hash func() hash.Hash
	}{
		{
			name: "ValidKey1",
			key:  "Some Key",
			iter: 4096,
			len:  32,
			hash: sha512.New,
		},
		{
			name: "ValidKey2",
			key:  "Another Key",
			iter: 8192,
			len:  64,
			hash: sha512.New,
		},
		{
			name: "ValidKey3",
			key:  "Yet Another Key",
			iter: 4096,
			len:  32,
			hash: sha512.New,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			someKeyKdf := kdf.PbKdf2.Key(tt.key, tt.iter, tt.len, tt.hash)
			bytes := someKeyKdf.Bytes()

			v, err := kdf.PbKdf2.FromBytes(bytes)
			if err != nil {
				t.Errorf("Failed to retrieve the value from bytes %s", err)
			} else if someKeyKdf.String() != v.String() {
				t.Errorf("Byte conversion failed original:\n\t%s\n\tdecoded:\n\t%s", someKeyKdf, v)
			}
		})
	}
}
