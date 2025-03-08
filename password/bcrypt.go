package password

import (
	"github.com/things-go/proc/internal/bytesconv"
	"golang.org/x/crypto/bcrypt"
)

var _ CryptFacade = Bcrypt{}

// Bcrypt bcrypt password encryption
type Bcrypt struct {
	cost int
}

func NewBcrypt(cost int) Bcrypt {
	if cost < bcrypt.MinCost || cost > bcrypt.MaxCost {
		cost = bcrypt.DefaultCost
	}
	return Bcrypt{cost: cost}
}

func (c Bcrypt) GenerateFromPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), c.cost)
	if err != nil {
		return "", err
	}
	return bytesconv.BytesToString(bytes), err
}

func (Bcrypt) CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword(bytesconv.StringToBytes(hashedPassword), bytesconv.StringToBytes(password))
}
