package password

import (
	"errors"
)

type CryptFacade interface {
	// GenerateFromPassword password hash encryption
	GenerateFromPassword(password string) (string, error)
	// CompareHashAndPassword password hash verification
	CompareHashAndPassword(hashedPassword, password string) error
}

const DefaultSaltSize = 16

// ErrCompareFailed compare failed
var ErrCompareFailed = errors.New("crypt compare failed")
