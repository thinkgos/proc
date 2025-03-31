package password

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"io"

	"github.com/thinkgos/proc/internal/bytesconv"
)

var _ CryptFacade = (*Simple)(nil)

// Simple simple password encryption
type Simple struct {
	saltSize int
}

// NewSimple new simple password encryption.
// if saltSize less or equal than zero, use default salt size.
func NewSimple(saltSize int) Simple {
	if saltSize <= 0 {
		saltSize = DefaultSaltSize
	}
	return Simple{saltSize: saltSize}
}

func (s Simple) GenerateFromPassword(password string) (string, error) {
	unencodedSalt := make([]byte, s.saltSize)
	_, err := io.ReadFull(rand.Reader, unencodedSalt)
	if err != nil {
		return "", err
	}
	pwd := s.hash(password, string(unencodedSalt))
	return base64.StdEncoding.EncodeToString(append(unencodedSalt, pwd...)), nil
}

func (s Simple) CompareHashAndPassword(hashedPassword, password string) error {
	orgRb, err := base64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return err
	}

	if len(orgRb) < s.saltSize {
		return ErrCompareFailed
	}
	unencodedSalt := orgRb[:s.saltSize]
	pwd := s.hash(password, bytesconv.BytesToString(unencodedSalt))
	if subtle.ConstantTimeCompare(orgRb[s.saltSize:], pwd) == 0 {
		return ErrCompareFailed
	}
	return nil
}

const (
	salt1 = `@#$%`
	salt2 = `^&*()`
)

func (Simple) hash(password, salt string) []byte {
	md5Pwd := md5.Sum(bytesconv.StringToBytes(password))

	build := &bytes.Buffer{}
	build.Grow(len(password) + len(salt)*3 + len(md5Pwd) + len(salt1) + len(salt2))
	// 加盐值加密
	build.WriteString(salt)
	build.WriteString(password)
	build.WriteString(salt1)
	build.WriteString(salt)
	build.Write(md5Pwd[:])
	build.WriteString(salt2)
	build.WriteString(salt)

	src := md5.Sum(build.Bytes())

	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src[:])
	return dst
}
