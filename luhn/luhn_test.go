package luhn_test

import (
	"testing"

	"github.com/thinkgos/proc/luhn"
)

// http://rosettacode.org/wiki/Luhn_test_of_credit_card_numbers
func Test_StdMod10(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{"49927398716", true},
		{"49927398717", false},
		{"1234567812345678", false},
		{"1234567812345670", true},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			if res := luhn.StdMod10.Validate(tc.input); res != tc.want {
				t.Errorf("Validate(%q) => %v, expected %v", tc.input, res, tc.want)
			}
		})
	}
}
