package luhn_test

import (
	"strings"
	"testing"

	"github.com/thinkgos/proc/luhn"
)

func Test_Unicode_BadAlphabet(t *testing.T) {
	var err error

	_, err = luhn.NewUnicode("")
	if err == nil {
		t.Fatal("expected error for empty alphabet")
	}
	_, err = luhn.NewUnicode("01234566789")
	if err == nil {
		t.Fatal("expected error for duplicate chars")
	}
	_, err = luhn.NewUnicode(strings.Repeat("abcdefghijklmn", 10))
	if err == nil {
		t.Fatal("expected error for too long alphabet")
	}
}

func Test_Unicode_Generate(t *testing.T) {
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
			a, err := luhn.NewUnicode(tc.chars)
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

func Test_Unicode_Validate(t *testing.T) {
	a, err := luhn.NewUnicode("abcdef")
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

// http://rosettacode.org/wiki/Luhn_test_of_credit_card_numbers
func Test_Unicode_ValidateRosetta(t *testing.T) {
	a, err := luhn.NewUnicode("0123456789")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

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
			if res := a.Validate(tc.input); res != tc.want {
				t.Errorf("Validate(%q) => %v, expected %v", tc.input, res, tc.want)
			}
		})
	}
}

// --- NewUnicode tests with real Unicode characters ---

func Test_NewUnicode_BadAlphabet(t *testing.T) {
	t.Run("invalid UTF-8", func(t *testing.T) {
		_, err := luhn.NewUnicode("\xc0\xaf")
		if err == nil {
			t.Fatal("expected error for invalid UTF-8")
		}
	})

	t.Run("empty string", func(t *testing.T) {
		_, err := luhn.NewUnicode("")
		if err == nil {
			t.Fatal("expected error for empty alphabet")
		}
	})

	t.Run("duplicate rune", func(t *testing.T) {
		_, err := luhn.NewUnicode("零一二三零")
		if err == nil {
			t.Fatal("expected error for duplicate rune")
		}
	})
}

func Test_NewUnicode_Valid(t *testing.T) {
	alphabets := []struct {
		name  string
		chars string
	}{
		{"CJK digits", "零一二三四五六七八九"},
		{"Katakana", "アイウエオカキクケコ"},
		{"Emoji", "😀😁😂🤣😃😄"},
		{"Cyrillic", "АБВГДЕЖЗИК"},
		{"Greek", "ΑΒΓΔΕΖΗΘΙΚ"},
		{"Mixed Latin+CJK", "AB一二三"},
	}
	for _, tc := range alphabets {
		t.Run(tc.name, func(t *testing.T) {
			l, err := luhn.NewUnicode(tc.chars)
			if err != nil {
				t.Fatalf("NewUnicode(%q) unexpected error: %v", tc.chars, err)
			}
			if l == nil {
				t.Fatal("expected non-nil LuhnModN")
			}
		})
	}
}

func Test_UnicodeModN_Generate(t *testing.T) {
	t.Run("CJK roundtrip", func(t *testing.T) {
		l, _ := luhn.NewUnicode("零一二三四五六七八九")
		input := "七九九二七三九八七一"
		check, err := l.Generate(input)
		if err != nil {
			t.Fatalf("Generate unexpected error: %v", err)
		}
		if len([]rune(check)) != 1 {
			t.Fatalf("expected single rune, got %q", check)
		}
		full := input + check
		if !l.Validate(full) {
			t.Errorf("Validate(%q) = false, want true", full)
		}
	})

	t.Run("emoji roundtrip", func(t *testing.T) {
		l, _ := luhn.NewUnicode("😀😁😂🤣😃😄")
		input := "😀😂😃🤣😁"
		check, err := l.Generate(input)
		if err != nil {
			t.Fatalf("Generate unexpected error: %v", err)
		}
		full := input + check
		if !l.Validate(full) {
			t.Errorf("Validate(%q) = false, want true", full)
		}
	})

	t.Run("katakana roundtrip", func(t *testing.T) {
		l, _ := luhn.NewUnicode("アイウエオカキクケコ")
		input := "カキクケ"
		check, err := l.Generate(input)
		if err != nil {
			t.Fatalf("Generate unexpected error: %v", err)
		}
		full := input + check
		if !l.Validate(full) {
			t.Errorf("Validate(%q) = false, want true", full)
		}
	})

	t.Run("mixed width roundtrip", func(t *testing.T) {
		l, _ := luhn.NewUnicode("AB一二三АБВ")
		input := "А一B二"
		check, err := l.Generate(input)
		if err != nil {
			t.Fatalf("Generate unexpected error: %v", err)
		}
		full := input + check
		if !l.Validate(full) {
			t.Errorf("Validate(%q) = false, want true", full)
		}
	})

	t.Run("invalid rune in input", func(t *testing.T) {
		l, _ := luhn.NewUnicode("😀😁😂🤣😃😄")
		_, err := l.Generate("😀😁X")
		if err == nil {
			t.Fatal("expected error for rune not in alphabet")
		}
	})

	t.Run("empty input", func(t *testing.T) {
		l, _ := luhn.NewUnicode("零一二三四五六七八九")
		_, err := l.Generate("")
		if err == nil {
			t.Fatal("expected error for empty input")
		}
	})
}

func Test_UnicodeModN_Validate(t *testing.T) {
	t.Run("CJK valid", func(t *testing.T) {
		l, _ := luhn.NewUnicode("零一二三四五六七八九")
		check, _ := l.Generate("七九九二七三九八七一")
		input := "七九九二七三九八七一" + check
		if !l.Validate(input) {
			t.Errorf("Validate(%q) = false, want true", input)
		}
	})

	t.Run("CJK wrong check digit", func(t *testing.T) {
		l, _ := luhn.NewUnicode("零一二三四五六七八九")
		// Append wrong check digit
		if l.Validate("七九九二七三九八七一零") {
			t.Error("expected wrong check digit to fail")
		}
	})

	t.Run("emoji valid", func(t *testing.T) {
		l, _ := luhn.NewUnicode("😀😁😂🤣😃😄")
		check, _ := l.Generate("😀😂😃🤣😁")
		input := "😀😂😃🤣😁" + check
		if !l.Validate(input) {
			t.Errorf("Validate(%q) = false, want true", input)
		}
	})

	t.Run("emoji tampered", func(t *testing.T) {
		l, _ := luhn.NewUnicode("😀😁😂🤣😃😄")
		if l.Validate("😀😂😃🤣😁😀") {
			t.Error("expected tampered emoji string to fail")
		}
	})

	t.Run("empty string", func(t *testing.T) {
		l, _ := luhn.NewUnicode("😀😁😂🤣😃😄")
		if l.Validate("") {
			t.Error("expected empty string to fail")
		}
	})

	t.Run("single rune", func(t *testing.T) {
		l, _ := luhn.NewUnicode("零一二三四五六七八九")
		// Single rune: Generate("零") should return a check digit that makes "零"+check valid
		check, err := l.Generate("零")
		if err != nil {
			t.Fatalf("Generate unexpected error: %v", err)
		}
		if !l.Validate("零" + check) {
			t.Errorf("Validate(\"零\"+check) = false, want true")
		}
	})
}

func Test_Unicode_Encode(t *testing.T) {
	testCases := []struct {
		name    string
		chars   string
		input   string
		wantErr bool
	}{
		{
			name:  "CJK",
			chars: "零一二三四五六七八九",
			input: "七九九二七三九八七一",
		},
		{
			name:  "Emoji",
			chars: "😀😁😂🤣😃😄",
			input: "😀😂😃🤣😁",
		},
		{
			name:  "Katakana",
			chars: "アイウエオカキクケコ",
			input: "カキクケ",
		},
		{
			name:  "Mixed width",
			chars: "AB一二三АБВ",
			input: "А一B二",
		},
		{
			name:    "Empty input",
			chars:   "零一二三四五六七八九",
			input:   "",
			wantErr: true,
		},
		{
			name:    "Invalid rune in input",
			chars:   "😀😁😂🤣😃😄",
			input:   "😀😁X",
			wantErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l, err := luhn.NewUnicode(tc.chars)
			if err != nil {
				t.Fatalf("NewUnicode(%q) unexpected error: %v", tc.chars, err)
			}
			encoded, err := l.Encode(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("Encode(%q) unexpected error: %v", tc.input, err)
			}
			// encoded should be input + one check rune
			runes := []rune(encoded)
			inputRunes := []rune(tc.input)
			if len(runes) != len(inputRunes)+1 {
				t.Errorf("Encode(%q) length = %d, want %d", tc.input, len(runes), len(inputRunes)+1)
			}
		})
	}
}

func Test_Unicode_Decode(t *testing.T) {
	t.Run("CJK valid", func(t *testing.T) {
		l, _ := luhn.NewUnicode("零一二三四五六七八九")
		payload := "七九九二七三九八七一"
		encoded, err := l.Encode(payload)
		if err != nil {
			t.Fatalf("Encode(%q) unexpected error: %v", payload, err)
		}
		decoded, err := l.Decode(encoded)
		if err != nil {
			t.Fatalf("Decode(%q) unexpected error: %v", encoded, err)
		}
		if decoded != payload {
			t.Errorf("Decode(%q) = %q, want %q", encoded, decoded, payload)
		}
	})

	t.Run("Emoji valid", func(t *testing.T) {
		l, _ := luhn.NewUnicode("😀😁😂🤣😃😄")
		payload := "😀😂😃🤣😁"
		encoded, err := l.Encode(payload)
		if err != nil {
			t.Fatalf("Encode(%q) unexpected error: %v", payload, err)
		}
		decoded, err := l.Decode(encoded)
		if err != nil {
			t.Fatalf("Decode(%q) unexpected error: %v", encoded, err)
		}
		if decoded != payload {
			t.Errorf("Decode(%q) = %q, want %q", encoded, decoded, payload)
		}
	})

	t.Run("Empty input", func(t *testing.T) {
		l, _ := luhn.NewUnicode("零一二三四五六七八九")
		_, err := l.Decode("")
		if err == nil {
			t.Fatal("expected error for empty input")
		}
	})

	t.Run("Wrong check digit", func(t *testing.T) {
		l, _ := luhn.NewUnicode("零一二三四五六七八九")
		// "零" is index 0, guaranteed to differ from any valid check digit for this payload
		if l.Validate("七九九二七三九八七一零") {
			t.Error("expected wrong check digit to fail")
		}
	})

	t.Run("Invalid rune in input", func(t *testing.T) {
		l, _ := luhn.NewUnicode("😀😁😂🤣😃😄")
		_, err := l.Decode("😀😁X")
		if err == nil {
			t.Fatal("expected error for rune not in alphabet")
		}
	})
}

func Test_Unicode_EncodeDecode_Roundtrip(t *testing.T) {
	alphabets := []struct {
		name  string
		chars string
		input string
	}{
		{"CJK", "零一二三四五六七八九", "七九九二七三九八七一"},
		{"Emoji", "😀😁😂🤣😃😄", "😀😂😃🤣😁"},
		{"Katakana", "アイウエオカキクケコ", "カキクケ"},
		{"Cyrillic", "АБВГДЕЖЗИК", "АБВГД"},
		{"Greek", "ΑΒΓΔΕΖΗΘΙΚ", "ΑΒΓΔ"},
		{"Mixed width", "AB一二三АБВ", "А一B二"},
		{"Single rune", "零一二三四五六七八九", "零"},
	}
	for _, tc := range alphabets {
		t.Run(tc.name, func(t *testing.T) {
			l, err := luhn.NewUnicode(tc.chars)
			if err != nil {
				t.Fatalf("NewUnicode(%q) unexpected error: %v", tc.chars, err)
			}
			encoded, err := l.Encode(tc.input)
			if err != nil {
				t.Fatalf("Encode(%q) unexpected error: %v", tc.input, err)
			}
			decoded, err := l.Decode(encoded)
			if err != nil {
				t.Fatalf("Decode(%q) unexpected error: %v", encoded, err)
			}
			if decoded != tc.input {
				t.Errorf("roundtrip: Encode(%q)=%q, Decode(%q)=%q, want %q", tc.input, encoded, encoded, decoded, tc.input)
			}
		})
	}
}
