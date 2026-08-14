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
	return Unicode{
		chars: []rune(s),
		index: seen,
	}, nil
}

func (l Unicode) Generate(s string) (string, error) {
	count := utf8.RuneCountInString(s)
	if count == 0 {
		return "", fmt.Errorf("luhn: input string must not be empty")
	}

	n := len(l.chars)
	factor := 1 + len(s)%2
	sum := 0
	for _, r := range s {
		codePoint, ok := l.index[r]
		if !ok {
			return "", fmt.Errorf("luhn: rune %q not valid in alphabet %q", r, string(l.chars))
		}
		addend := factor * codePoint
		factor = 3 - factor
		addend = (addend / n) + (addend % n)
		sum += addend
	}

	remainder := sum % n
	checkCodePoint := (n - remainder) % n
	return string(l.chars[checkCodePoint]), nil
}

func (l Unicode) Validate(s string) bool {
	runes := []rune(s)
	if len(runes) == 0 {
		return false
	}
	s1, c1 := string(runes[:len(runes)-1]), runes[len(runes)-1]
	c2, err := l.Generate(s1)
	return err == nil && string(c1) == c2
}
