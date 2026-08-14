package luhn

import (
	"errors"
	"fmt"
)

// An Alphabet is a string of N Alphabet characters, representing the digits of a given base N.
type Alphabet struct {
	chars []byte
	// index maps a byte to its position in chars. -1 means "not in alphabet".
	index [128]int8
}

// New converts the given string into an Alphabet, verifying that it is
// composed of unique ascii characters.
func New(s string) (LuhnModN, error) {
	if len(s) == 0 {
		return nil, errors.New("luhn: alphabet must not be empty")
	}
	if len(s) > 127 {
		return nil, errors.New("luhn: alphabet too large")
	}

	var index [128]int8
	for i := range index {
		index[i] = -1
	}

	chars := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 128 {
			return nil, fmt.Errorf("luhn: character %q at position %d is not ASCII", c, i)
		}
		if index[c] != -1 {
			return nil, fmt.Errorf("luhn: character %q non-unique in alphabet %q", c, s)
		}
		index[c] = int8(i)
		chars[i] = c
	}
	return Alphabet{
		chars: chars,
		index: index,
	}, nil
}

func (l Alphabet) Generate(s string) (string, error) {
	if len(s) == 0 {
		return "", fmt.Errorf("luhn: input string must not be empty")
	}
	n := len(l.chars)
	factor := 1 + len(s)%2
	sum := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 128 {
			return "", fmt.Errorf("luhn: character %q at position %d is not ASCII", c, i)
		}
		codePoint := int(l.index[c])
		if codePoint == -1 {
			return "", fmt.Errorf("luhn: character %q not valid in alphabet %q", c, string(l.chars))
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

func (l Alphabet) Validate(s string) bool {
	if len(s) == 0 {
		return false
	}
	t, last := s[:len(s)-1], s[len(s)-1:]
	c, err := l.Generate(t)
	return err == nil && last == c
}
