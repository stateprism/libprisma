package pkcrypto

import (
	"crypto"
	"hash"
)

type PKCrypto interface {
	NewKey(options ...any) crypto.PrivateKey
	Sign(k crypto.PrivateKey, message []byte, h func() hash.Hash) ([]byte, error)
	Verify(k crypto.PublicKey, sig []byte, message []byte, opts ...any) (bool, error)
}
