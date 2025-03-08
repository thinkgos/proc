package password

import (
	"errors"
	"sync/atomic"
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

// 考虑有些场景是一些低性能机器, 在使用bcrypt时, 很耗时
var crypt = &atomic.Pointer[CryptFacade]{}

func init() {
	SetCrypt(NewBcrypt(0))
}

func SetCrypt(pf CryptFacade) {
	crypt.Store(&pf)
}

// GenerateFromPassword password hash encryption
func GenerateFromPassword(password string) (string, error) {
	return (*crypt.Load()).GenerateFromPassword(password)
}

// CompareHashAndPassword password hash verification
func CompareHashAndPassword(hashedPassword, password string) error {
	return (*crypt.Load()).CompareHashAndPassword(hashedPassword, password)
}
