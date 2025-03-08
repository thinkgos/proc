// Copyright [2020] [thinkgos] thinkgo@aliyun.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"io"

	"github.com/things-go/proc/internal/bytesconv"
	"golang.org/x/crypto/scrypt"
)

var _ CryptFacade = Scrypt{}

// Scrypt scrypt password encryption
type Scrypt struct {
	saltSize int
}

// NewScrypt new scrypt password encryption.
// if saltSize less or equal than zero, use default salt size.
func NewScrypt(saltSize int) *Scrypt {
	if saltSize <= 0 {
		saltSize = DefaultSaltSize
	}
	return &Scrypt{saltSize: saltSize}
}

func (s Scrypt) GenerateFromPassword(password string) (string, error) {
	unencodedSalt := make([]byte, s.saltSize)
	_, err := io.ReadFull(rand.Reader, unencodedSalt)
	if err != nil {
		return "", err
	}

	rb, err := scrypt.Key(bytesconv.StringToBytes(password), unencodedSalt, 16384, 8, 1, 32)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(append(unencodedSalt, rb...)), nil
}

func (s Scrypt) CompareHashAndPassword(hashedPassword, password string) error {
	orgRb, err := base64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return err
	}
	if len(orgRb) < s.saltSize {
		return ErrCompareFailed
	}
	unencodedSalt := orgRb[:s.saltSize]
	rb, err := scrypt.Key(bytesconv.StringToBytes(password), unencodedSalt, 16384, 8, 1, 32)
	if err != nil {
		return err
	}
	if subtle.ConstantTimeCompare(orgRb[s.saltSize:], rb) == 0 {
		return ErrCompareFailed
	}
	return nil
}
