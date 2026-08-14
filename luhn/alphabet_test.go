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

func Test_Encode(t *testing.T) {
	testCases := []struct {
		name    string
		chars   string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:  "Base 10 credit card",
			chars: "0123456789",
			input: "7992739871",
			want:  "79927398713",
		},
		{
			name:  "Base 6",
			chars: "abcdef",
			input: "abcdef",
			want:  "abcdefe",
		},
		{
			name:  "Base 32",
			chars: "234567ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			input: "ABCDEF",
			want:  "ABCDEFM",
		},
		{
			name:  "Single char",
			chars: "0123456789",
			input: "7",
			want:  "75",
		},
		{
			name:    "Empty input",
			chars:   "0123456789",
			input:   "",
			wantErr: true,
		},
		{
			name:    "Invalid char in input",
			chars:   "0123456789",
			input:   "7992739871x",
			wantErr: true,
		},
		{
			name:    "Non-ASCII char in input",
			chars:   "0123456789",
			input:   "799中",
			wantErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a, err := luhn.New(tc.chars)
			if err != nil {
				t.Fatalf("New(%q) unexpected error: %v", tc.chars, err)
			}
			got, err := a.Encode(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("Encode(%q) unexpected error: %v", tc.input, err)
			}
			if got != tc.want {
				t.Errorf("Encode(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func Test_Decode(t *testing.T) {
	testCases := []struct {
		name    string
		chars   string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:  "Base 10 credit card valid",
			chars: "0123456789",
			input: "79927398713",
			want:  "7992739871",
		},
		{
			name:  "Base 6 valid",
			chars: "abcdef",
			input: "abcdeb",
			want:  "abcde",
		},
		{
			name:  "Base 32 valid",
			chars: "234567ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			input: "ABCDEFM",
			want:  "ABCDEF",
		},
		{
			name:    "Empty input",
			chars:   "0123456789",
			input:   "",
			wantErr: true,
		},
		{
			name:    "Wrong check digit",
			chars:   "0123456789",
			input:   "79927398710",
			wantErr: true,
		},
		{
			name:  "Two chars (single payload + check)",
			chars: "0123456789",
			input: "75",
			want:  "7",
		},
		{
			name:    "Invalid char in input",
			chars:   "0123456789",
			input:   "7992739871x",
			wantErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a, err := luhn.New(tc.chars)
			if err != nil {
				t.Fatalf("New(%q) unexpected error: %v", tc.chars, err)
			}
			got, err := a.Decode(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("Decode(%q) unexpected error: %v", tc.input, err)
			}
			if got != tc.want {
				t.Errorf("Decode(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func Test_EncodeDecode_Roundtrip(t *testing.T) {
	alphabets := []struct {
		name  string
		chars string
		input string
	}{
		{"Mod10", "0123456789", "7992739871"},
		{"Mod6", "abcdef", "abcdef"},
		{"Mod32", "234567ABCDEFGHIJKLMNOPQRSTUVWXYZ", "ABCDEFGH"},
		{"Mod36", "0123456789abcdefghijklmnopqrstuvwxyz", "hello123"},
		{"Mod62", "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "abcXYZ123"},
		{"Single char", "0123456789", "5"},
	}
	for _, tc := range alphabets {
		t.Run(tc.name, func(t *testing.T) {
			a, err := luhn.New(tc.chars)
			if err != nil {
				t.Fatalf("New(%q) unexpected error: %v", tc.chars, err)
			}
			encoded, err := a.Encode(tc.input)
			if err != nil {
				t.Fatalf("Encode(%q) unexpected error: %v", tc.input, err)
			}
			decoded, err := a.Decode(encoded)
			if err != nil {
				t.Fatalf("Decode(%q) unexpected error: %v", encoded, err)
			}
			if decoded != tc.input {
				t.Errorf("roundtrip: Encode(%q)=%q, Decode(%q)=%q, want %q", tc.input, encoded, encoded, decoded, tc.input)
			}
		})
	}
}
