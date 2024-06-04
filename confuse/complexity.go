// Copyright [2022] [thinkgos] thinkgo@aliyun.com
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

package confuse

import (
	cryptorand "crypto/rand"
	"math/big"
	"math/rand/v2"
	"strings"

	"github.com/things-go/proc/internal/bytesconv"
)

const (
	lower = `abcdefghijklmnopqrstuvwxyz`
	upper = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	digit = `0123456789`
	spec  = ` !"#$%&'()*+,-./:;<=>?@[\]^_{|}~` + "`"
)

type complexityConfig struct {
	lower bool
	upper bool
	digit bool
	spec  bool
	meet  bool
}

// Option for Complexity
type Option func(*complexityConfig)

// WithLower use lower
func WithLower() Option {
	return func(c *complexityConfig) {
		c.lower = true
	}
}

// WithUpper use upper
func WithUpper() Option {
	return func(c *complexityConfig) {
		c.upper = true
	}
}

// WithDigit use digit
func WithDigit() Option {
	return func(c *complexityConfig) {
		c.digit = true
	}
}

// WithLowerUpperDigit use lower upper digit
func WithLowerUpperDigit() Option {
	return func(c *complexityConfig) {
		c.lower = true
		c.upper = true
		c.digit = true
	}
}

// WithLowerUpper use lower upper
func WithLowerUpper() Option {
	return func(c *complexityConfig) {
		c.lower = true
		c.upper = true
	}
}

// WithSpec use spec
func WithSpec() Option {
	return func(c *complexityConfig) {
		c.spec = true
	}
}

// WithAll use lower upper digit spec and enable meet complexity
func WithAll() Option {
	return func(c *complexityConfig) {
		c.lower = true
		c.upper = true
		c.digit = true
		c.spec = true
		c.meet = true
	}
}

// WithMeet enable meet complexity
func WithMeet() Option {
	return func(c *complexityConfig) {
		c.meet = true
	}
}

// Complexity setting
type Complexity struct {
	chars        string
	requiredList []string
}

// NewComplexity new complexity with option
// default use lower, upper and digit to generate a random string, and not meets complexity.
func NewComplexity(opts ...Option) *Complexity {
	c := complexityConfig{false, false, false, false, false}
	for _, opt := range opts {
		opt(&c)
	}

	co := &Complexity{}
	if c.lower {
		co.chars += lower
		co.requiredList = append(co.requiredList, lower)
	}
	if c.upper {
		co.chars += upper
		co.requiredList = append(co.requiredList, upper)
	}
	if c.digit {
		co.chars += digit
		co.requiredList = append(co.requiredList, digit)
	}
	if c.spec {
		co.chars += spec
		co.requiredList = append(co.requiredList, spec)
	}

	if c.meet {
		if co.chars == "" {
			co.requiredList = []string{lower, upper, digit}
		}
	} else {
		co.requiredList = nil
	}
	if co.chars == "" {
		co.chars = lower + upper + digit
	}
	return co
}

// IsComplexEnough return True if s meets complexity settings
func (sf *Complexity) IsComplexEnough(s string) bool {
	for _, chars := range sf.requiredList {
		if !strings.ContainsAny(chars, s) {
			return false
		}
	}
	return true
}

// Generate a random string which is complex enough.
func (sf *Complexity) Generate(n int) string {
	var idx int

	buffer := make([]byte, n)
	max := big.NewInt(int64(len(sf.chars)))
	for {
		for j := 0; j < n; j++ {
			rnd, err := cryptorand.Int(cryptorand.Reader, max)
			if err != nil {
				idx = rand.IntN(len(sf.chars))
			} else {
				idx = int(rnd.Int64())
			}
			buffer[j] = sf.chars[idx]
		}
		v := bytesconv.BytesToString(buffer)
		if sf.IsComplexEnough(v) && v[:1] != " " && v[n-1:] != " " {
			return v
		}
	}
}

// use lower, upper and digit to generate a random string, and meets defaultComplexity.
var defaultComplexity = NewComplexity(WithMeet())

// IsComplexEnough return True if s meets complexity settings.
// which use lower, upper and digit, and meets complexity.
func IsComplexEnough(s string) bool {
	return defaultComplexity.IsComplexEnough(s)
}

// Generate a random string which is complex enough.
// which use lower, upper and digit, and meets complexity.
func Generate(n int) string {
	return defaultComplexity.Generate(n)
}
