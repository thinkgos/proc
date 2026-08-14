package luhn

// https://en.wikipedia.org/wiki/ISO/IEC_7064
type LuhnModN interface {
	// Generate returns a check digit for the string s, which should be composed
	// of characters from the Alphabet.
	Generate(s string) (string, error)
	// Validate returns true if the last character of the string s is correct, for
	// a string s composed of characters in the alphabet.
	Validate(s string) bool
}

// https://en.wikipedia.org/wiki/Luhn_algorithm
// http://rosettacode.org/wiki/Luhn_test_of_credit_card_numbers
// Use "0123456789"
var StdMod10 LuhnModN

// Use "234567ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var StdMod32 LuhnModN

// Use "0123456789ABCDEFGHIJKLMNOPQRSTUV"
var StdMod32Hex LuhnModN

// Use "0123456789abcdefghijklmnopqrstuvwxyz"
var StdMod36 LuhnModN

// Use "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
var StdMod58 LuhnModN

func init() {
	var err error

	StdMod10, err = New("0123456789")
	if err != nil {
		panic(err)
	}
	StdMod32, err = New("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567")
	if err != nil {
		panic(err)
	}
	StdMod32Hex, err = New("0123456789ABCDEFGHIJKLMNOPQRSTUV")
	if err != nil {
		panic(err)
	}
	StdMod36, err = New("0123456789abcdefghijklmnopqrstuvwxyz")
	if err != nil {
		panic(err)
	}
	StdMod58, err = New("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")
	if err != nil {
		panic(err)
	}
}
