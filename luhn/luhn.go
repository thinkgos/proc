package luhn

// https://en.wikipedia.org/wiki/ISO/IEC_7064
type LuhnModN interface {
	// Generate returns a check digit for the string s, which should be composed
	// of characters from the Alphabet.
	Generate(s string) (string, error)
	// Validate returns true if the last character of the string s is correct, for
	// a string s composed of characters in the alphabet.
	Validate(s string) bool
	// Encode appends a Luhn check character to the end of the input string s.
	Encode(s string) (string, error)
	// Decode validates the input string s and returns the original string
	// with the trailing check character removed. It returns an error if
	// the string is empty or fails validation.
	Decode(s string) (string, error)
}

// https://en.wikipedia.org/wiki/Luhn_algorithm
// http://rosettacode.org/wiki/Luhn_test_of_credit_card_numbers
// Use "0123456789"
var StdMod10 *Alphabet

// Use "234567ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var StdMod32 *Alphabet

// Use "0123456789ABCDEFGHIJKLMNOPQRSTUV"
var StdMod32Hex *Alphabet

// Use "0123456789abcdefghijklmnopqrstuvwxyz"
var StdMod36 *Alphabet

// Use "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
var StdMod58 *Alphabet

// Use "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var StdMod62 *Alphabet

func init() {
	var err error

	StdMod10, err = newAlphabet("0123456789")
	if err != nil {
		panic(err)
	}
	StdMod32, err = newAlphabet("234567ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if err != nil {
		panic(err)
	}
	StdMod32Hex, err = newAlphabet("0123456789ABCDEFGHIJKLMNOPQRSTUV")
	if err != nil {
		panic(err)
	}
	StdMod36, err = newAlphabet("0123456789abcdefghijklmnopqrstuvwxyz")
	if err != nil {
		panic(err)
	}
	StdMod58, err = newAlphabet("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")
	if err != nil {
		panic(err)
	}
	StdMod62, err = newAlphabet("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if err != nil {
		panic(err)
	}
}
