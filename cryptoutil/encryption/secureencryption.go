package encryption

import "errors"

type CypherError int

const (
	ErrInvalidKeyLength CypherError = iota
	ErrTagMismatch      CypherError = iota
	ErrorUnknown        CypherError = iota
)

func (e CypherError) String() string {
	switch {
	case errors.Is(e, ErrInvalidKeyLength):
		return "ErrInvalidKeyLength"
	case errors.Is(e, ErrTagMismatch):
		return "ErrTagMismatch"
	case errors.Is(e, ErrorUnknown):
		return "ErrorUnknown"
	default:
		return "ErrorUnknown"
	}
}

func (e CypherError) Error() string {
	return e.String()
}

type SecureCypher interface {
	Encrypt(data []byte) ([]byte, error)
	Decrypt(data []byte) ([]byte, error)
	GetTag() []byte
	GetTagSize() int
	GetIV() []byte
	SetIV(iv []byte)
	GetIvSize() int
	TagPlusIVSize() int
	GetKeyThumbprint() []byte
	CheckKeyThumbprint(thumbprint []byte) bool
	Reset()
	FullReset()
	Dispose()
	EncryptToBytes(data []byte) ([]byte, error)
	DecryptFromBytes(data []byte) ([]byte, error)
}
