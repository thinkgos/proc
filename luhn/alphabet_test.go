package luhn_test

import (
	"strings"
	"testing"

	"github.com/thinkgos/proc/luhn"
)

func Test_BadAlphabet(t *testing.T) {
	var err error

	_, err = luhn.New("")
	if err == nil {
		t.Fatal("expected error for empty alphabet")
	}
	_, err = luhn.New("01234566789")
	if err == nil {
		t.Fatal("expected error for duplicate chars")
	}
	_, err = luhn.New(strings.Repeat("abcdefghijklmn", 10))
	if err == nil {
		t.Fatal("expected error for too long alphabet")
	}
	_, err = luhn.New("中")
	if err == nil {
		t.Fatal("expected error for non-ASCII alphabet")
	}
}

func Test_Generate(t *testing.T) {
	testCases := []struct {
		name    string
		chars   string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "Base 6",
			chars:   "abcdef",
			input:   "abcdef",
			want:    "e",
			wantErr: false,
		},
		{
			name:    "Base 10",
			chars:   "0123456789",
			input:   "7992739871",
			want:    "3",
			wantErr: false,
		},
		{
			name:    "Invalid: input not valid in alphabet",
			chars:   "ABC",
			input:   "7992739871",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Invalid: empty input string",
			chars:   "ABC",
			input:   "",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Invalid: not valid input string",
			chars:   "ABC",
			input:   "ABC中",
			want:    "",
			wantErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.chars+" "+tc.input, func(t *testing.T) {
			a, err := luhn.New(tc.chars)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			res, err := a.Generate(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if res != tc.want {
					t.Errorf("Generate(%q) = %q, want %q", tc.input, res, tc.want)
				}
			}
		})
	}
}

func Test_Validate(t *testing.T) {
	a, err := luhn.New("abcdef")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	testCases := []struct {
		input string
		want  bool
	}{
		{
			input: "abcdefe",
			want:  true,
		},
		{
			input: "abcdefd",
			want:  false,
		},
		{
			input: "",
			want:  false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			if res := a.Validate(tc.input); res != tc.want {
				t.Errorf("Validate(%q) => %v, expected %v", tc.input, res, tc.want)
			}
		})
	}
}
