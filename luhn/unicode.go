package luhn

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// An Unicode is a string of N characters, representing the digits of a given base N.
type Unicode struct {
	chars []rune
	// index maps a rune to its position in chars.
	index map[rune]int
}

// NewUnicode converts the given string an Alphabet, verifying that it is correct.
func NewUnicode(s string) (LuhnModN, error) {
	return newUnicode(s)
}

// NewUnicode converts the given string an Alphabet, verifying that it is correct.
func newUnicode(s string) (*Unicode, error) {
	if !utf8.ValidString(s) {
		return nil, errors.New("luhn: alphabet contains invalid UTF-8")
	}
	count := utf8.RuneCountInString(s)
	if count == 0 {
		return nil, errors.New("luhn: alphabet must not be empty")
	}
	pos := 0
	seen := make(map[rune]int, count)
	for _, r := range s {
		if _, ok := seen[r]; ok {
			return nil, fmt.Errorf("luhn: character %q non-unique in alphabet %q", r, s)
		}
		seen[r] = pos
		pos++
	}
	return &Unicode{
		chars: []rune(s),
		index: seen,
	}, nil
}

// Generate returns a check digit for the string s, which should be composed
// of characters from the Alphabet.
func (l Unicode) Generate(s string) (string, error) {
	b, err := l.generate(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Validate returns true if the last character of the string s is correct, for
// a string s composed of characters in the alphabet.
func (l Unicode) Validate(s string) bool {
	runes := []rune(s)
	if len(runes) < 1 {
		return false
	}
	s1, last := string(runes[:len(runes)-1]), runes[len(runes)-1]
	c, err := l.generate(s1)
	return err == nil && last == c
}

// Encode appends a Luhn check character to the end of the input string s.
func (l Unicode) Encode(s string) (string, error) {
	checkRune, err := l.generate(s)
	if err != nil {
		return "", err
	}
	return s + string(checkRune), nil
}

// Decode validates the input string s and returns the original string
// with the trailing check character removed. It returns an error if
// the string is empty, contains invalid UTF-8, or fails validation.
func (l Unicode) Decode(s string) (string, error) {
	runes := []rune(s)
	if len(runes) == 0 {
		return "", errors.New("luhn: alphabet must not be empty")
	}
	s1, last := string(runes[:len(runes)-1]), runes[len(runes)-1]
	c, err := l.generate(s1)
	if err == nil && last == c {
		return s1, nil
	}
	return "", errors.New("luhn: checksum verification failed")
}

func (l Unicode) generate(s string) (rune, error) {
	count := utf8.RuneCountInString(s)
	if count == 0 {
		return 0, fmt.Errorf("luhn: input string must not be empty")
	}

	n := len(l.chars)
	factor := 1 + len(s)%2
	sum := 0
	for _, r := range s {
		codePoint, ok := l.index[r]
		if !ok {
			return 0, fmt.Errorf("luhn: rune %q not valid in alphabet %q", r, string(l.chars))
		}
		addend := factor * codePoint
		factor = 3 - factor
		addend = (addend / n) + (addend % n)
		sum += addend
	}

	remainder := sum % n
	checkCodePoint := (n - remainder) % n
	return l.chars[checkCodePoint], nil
}
