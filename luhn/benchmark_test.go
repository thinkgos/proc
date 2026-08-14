package luhn_test

import (
	"testing"

	"github.com/thinkgos/proc/luhn"
)

// BenchmarkAlphabetGenerate benchmarks Alphabet.Generate
func BenchmarkAlphabetGenerate(b *testing.B) {
	benchmarks := []struct {
		name  string
		l     luhn.LuhnModN
		input string
	}{
		{"Mod10_Short", luhn.StdMod10, "4992739871"},
		{"Mod10_Medium", luhn.StdMod10, "499273987161234"},
		{"Mod10_Long", luhn.StdMod10, "49927398716123456789012345"},
		{"Mod32_Short", luhn.StdMod32, "ABCDEF"},
		{"Mod62_Short", luhn.StdMod62, "abc123DEF"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for b.Loop() {
				_, _ = bm.l.Generate(bm.input)
			}
		})
	}
}

// BenchmarkAlphabetValidate benchmarks Alphabet.Validate
func BenchmarkAlphabetValidate(b *testing.B) {
	benchmarks := []struct {
		name  string
		l     luhn.LuhnModN
		input string
	}{
		{"Mod10_Valid", luhn.StdMod10, "49927398716"},
		{"Mod10_Invalid", luhn.StdMod10, "49927398717"},
		{"Mod32_Valid", luhn.StdMod32, "ABCDEFG3"},
		{"Mod62_Valid", luhn.StdMod62, "abc123DEF5"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for b.Loop() {
				_ = bm.l.Validate(bm.input)
			}
		})
	}
}

// BenchmarkUnicodeGenerate benchmarks Unicode.Generate
func BenchmarkUnicodeGenerate(b *testing.B) {
	unicodeMod10, _ := luhn.NewUnicode("0123456789")
	unicodeMod36, _ := luhn.NewUnicode("0123456789abcdefghijklmnopqrstuvwxyz")

	benchmarks := []struct {
		name  string
		l     luhn.LuhnModN
		input string
	}{
		{"Unicode_Mod10_Short", unicodeMod10, "4992739871"},
		{"Unicode_Mod10_Medium", unicodeMod10, "499273987161234"},
		{"Unicode_Mod36_Short", unicodeMod36, "abc123"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for b.Loop() {
				_, _ = bm.l.Generate(bm.input)
			}
		})
	}
}

// BenchmarkUnicodeValidate benchmarks Unicode.Validate
func BenchmarkUnicodeValidate(b *testing.B) {
	unicodeMod10, _ := luhn.NewUnicode("0123456789")

	benchmarks := []struct {
		name  string
		l     luhn.LuhnModN
		input string
	}{
		{"Unicode_Mod10_Valid", unicodeMod10, "49927398716"},
		{"Unicode_Mod10_Invalid", unicodeMod10, "49927398717"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for b.Loop() {
				_ = bm.l.Validate(bm.input)
			}
		})
	}
}

// BenchmarkAlphabetEncode benchmarks Alphabet.Encode
func BenchmarkAlphabetEncode(b *testing.B) {
	benchmarks := []struct {
		name  string
		l     luhn.LuhnModN
		input string
	}{
		{"Mod10_Short", luhn.StdMod10, "4992739871"},
		{"Mod10_Medium", luhn.StdMod10, "499273987161234"},
		{"Mod10_Long", luhn.StdMod10, "49927398716123456789012345"},
		{"Mod32_Short", luhn.StdMod32, "ABCDEF"},
		{"Mod62_Short", luhn.StdMod62, "abc123DEF"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for b.Loop() {
				_, _ = bm.l.Encode(bm.input)
			}
		})
	}
}

// BenchmarkAlphabetDecode benchmarks Alphabet.Decode
func BenchmarkAlphabetDecode(b *testing.B) {
	benchmarks := []struct {
		name  string
		l     luhn.LuhnModN
		input string
	}{
		{"Mod10_Short", luhn.StdMod10, "49927398716"},
		{"Mod10_Medium", luhn.StdMod10, "4992739871612345"},
		{"Mod10_Long", luhn.StdMod10, "499273987161234567890123451"},
		{"Mod32_Short", luhn.StdMod32, "ABCDEFG3"},
		{"Mod62_Short", luhn.StdMod62, "abc123DEF5"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for b.Loop() {
				_, _ = bm.l.Decode(bm.input)
			}
		})
	}
}

// BenchmarkUnicodeEncode benchmarks Unicode.Encode
func BenchmarkUnicodeEncode(b *testing.B) {
	unicodeMod10, _ := luhn.NewUnicode("0123456789")
	unicodeCJK, _ := luhn.NewUnicode("零一二三四五六七八九")

	benchmarks := []struct {
		name  string
		l     luhn.LuhnModN
		input string
	}{
		{"Unicode_Mod10_Short", unicodeMod10, "4992739871"},
		{"Unicode_Mod10_Medium", unicodeMod10, "499273987161234"},
		{"Unicode_CJK_Short", unicodeCJK, "七九九二七三"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for b.Loop() {
				_, _ = bm.l.Encode(bm.input)
			}
		})
	}
}

// BenchmarkUnicodeDecode benchmarks Unicode.Decode
func BenchmarkUnicodeDecode(b *testing.B) {
	unicodeMod10, _ := luhn.NewUnicode("0123456789")
	unicodeCJK, _ := luhn.NewUnicode("零一二三四五六七八九")

	benchmarks := []struct {
		name  string
		l     luhn.LuhnModN
		input string
	}{
		{"Unicode_Mod10_Short", unicodeMod10, "49927398716"},
		{"Unicode_Mod10_Medium", unicodeMod10, "4992739871612345"},
		{"Unicode_CJK_Short", unicodeCJK, "七九九二七三三"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for b.Loop() {
				_, _ = bm.l.Decode(bm.input)
			}
		})
	}
}

// BenchmarkStdMods benchmarks all standard Mod implementations
func BenchmarkStdMods(b *testing.B) {
	type modTest struct {
		name  string
		l     luhn.LuhnModN
		input string
	}
	tests := []modTest{
		{"StdMod10", luhn.StdMod10, "49927398716"},
		{"StdMod32", luhn.StdMod32, "ABCDEFG3"},
		{"StdMod32Hex", luhn.StdMod32Hex, "0123456789ABCDEF5"},
		{"StdMod36", luhn.StdMod36, "abc123def4"},
		{"StdMod58", luhn.StdMod58, "abc123XYZ"},
		{"StdMod62", luhn.StdMod62, "abc123XYZ789"},
	}

	b.Run("Generate", func(b *testing.B) {
		for _, tt := range tests {
			b.Run(tt.name, func(b *testing.B) {
				for b.Loop() {
					_, _ = tt.l.Generate(tt.input)
				}
			})
		}
	})

	b.Run("Validate", func(b *testing.B) {
		for _, tt := range tests {
			b.Run(tt.name, func(b *testing.B) {
				for b.Loop() {
					_ = tt.l.Validate(tt.input)
				}
			})
		}
	})

	b.Run("Encode", func(b *testing.B) {
		for _, tt := range tests {
			// Encode expects input without check digit, trim the last character
			input := tt.input[:len(tt.input)-1]
			b.Run(tt.name, func(b *testing.B) {
				for b.Loop() {
					_, _ = tt.l.Encode(input)
				}
			})
		}
	})

	b.Run("Decode", func(b *testing.B) {
		for _, tt := range tests {
			b.Run(tt.name, func(b *testing.B) {
				for b.Loop() {
					_, _ = tt.l.Decode(tt.input)
				}
			})
		}
	})
}

// BenchmarkNew benchmarks New function for creating Alphabet
func BenchmarkNew(b *testing.B) {
	for b.Loop() {
		_, _ = luhn.New("0123456789")
	}
}

// BenchmarkNewUnicode benchmarks NewUnicode function for creating Unicode
func BenchmarkNewUnicode(b *testing.B) {
	for b.Loop() {
		_, _ = luhn.NewUnicode("0123456789")
	}
}
