package pkcrypto

import (
	"crypto"
	"fmt"
	"github.com/stateprism/libprisma/cryptoutil"
	"golang.org/x/crypto/ed25519"
	"hash"
)

type ed25519Impl struct{}

var Ed25519 ed25519Impl

func (ed25519Impl) NewKey(...any) crypto.PrivateKey {
	return ed25519.NewKeyFromSeed(cryptoutil.NewRandom(32))
}

func (ed25519Impl) Sign(k crypto.PrivateKey, message []byte, h func() hash.Hash) ([]byte, error) {
	hx := h()
	hd := make([]byte, 0)
	if n, err := hx.Write(message); err != nil {
		return nil, err
	} else if n != len(message) {
		return nil, fmt.Errorf("hasher read %d and not the expected %d", n, len(message))
	}
	hd = hx.Sum([]byte{})
	return ed25519.Sign(k.(ed25519.PrivateKey), hd), nil
}

func (ed25519Impl) Verify(k crypto.PublicKey, sig []byte, message []byte, opts ...any) (bool, error) {
	return ed25519.Verify(k.(ed25519.PublicKey), message, sig), nil
}
