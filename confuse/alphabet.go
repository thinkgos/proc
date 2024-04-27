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
	cryptoRand "crypto/rand"
	"math/bits"
	"math/rand/v2"
	"unsafe"
)

// Previous defined bytes, do not change this.
var (
	DefaultAlphabet   = []byte("QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz")
	DefaultDigit      = []byte("0123456789")
	DefaultAlphaDigit = []byte("QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz0123456789")
	DefaultSymbol     = []byte("QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_{|}~`") //nolint: lll
)

// Alphabet rand alpha with give length, which Contains only letters
func Alphabet(length int) string { return randString(length, DefaultAlphabet) }

// AlphabetBytes rand alpha with give length, which Contains only letters
func AlphabetBytes(length int) []byte { return randBytes(length, DefaultAlphabet) }

// Number rand string with give length, which Contains only number
func Number(length int) string { return randString(length, DefaultDigit) }

// NumberBytes rand string with give length, which Contains only number
func NumberBytes(length int) []byte { return randBytes(length, DefaultDigit) }

// AlphaNumber rand string with give length, which Contains number and letters
func AlphaNumber(length int) string { return randString(length, DefaultAlphaDigit) }

// AlphaNumberBytes rand string with give length, which Contains number and letters
func AlphaNumberBytes(length int) []byte { return randBytes(length, DefaultAlphaDigit) }

// Symbol rand symbol with give length, which Contains number, letters and specific symbol
func Symbol(length int) string { return randString(length, DefaultSymbol) }

// SymbolBytes rand symbol with give length, which Contains number, letters and specific symbol
func SymbolBytes(length int) []byte { return randBytes(length, DefaultSymbol) }

// String rand bytes, if not alphabets, it will use DefaultAlphabet.
func String(length int, alphabets ...byte) string {
	if len(alphabets) == 0 {
		alphabets = DefaultAlphaDigit
	}
	return randString(length, alphabets)
}

// Bytes rand bytes, if not alphabets, it will use DefaultAlphabet.
func Bytes(length int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = DefaultAlphaDigit
	}
	return randBytes(length, alphabets)
}

func randString(length int, alphabets []byte) string {
	b := randBytes(length, alphabets)
	return *(*string)(unsafe.Pointer(&b))
}

func randBytes(length int, alphabets []byte) []byte {
	b := make([]byte, length)
	if _, err := cryptoRand.Read(b); err == nil {
		for i, v := range b {
			b[i] = alphabets[v%byte(len(alphabets))]
		}
		return b
	}

	bn := bits.Len(uint(len(alphabets)))
	mask := int64(1)<<bn - 1
	max := 63 / bn

	// A rand.Int64() generates 63 random bits, enough for alphabets letters!
	for i, cache, remain := 0, rand.Int64(), max; i < length; {
		if remain == 0 {
			cache, remain = rand.Int64(), max
		}
		if idx := int(cache & mask); idx < len(alphabets) {
			b[i] = alphabets[idx]
			i++
		}
		cache >>= bn
		remain--
	}
	return b
}
