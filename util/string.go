package util

import "unicode/utf8"

// Substr : subtract string
func Substr(s string, start, end int) string {
	b := []byte(s)
	idx1 := 0
	idx2 := 0

	for i := 0; i < start; i++ {
		_, size := utf8.DecodeRune(b[idx1:])
		idx1 += size
	}

	for i := 0; i < end; i++ {
		_, size := utf8.DecodeRune(b[idx2:])
		idx2 += size
	}

	return s[idx1:idx2]
}

// GetCharStr : get char
func GetCharStr(s string, sIdx int) string {
	b := []byte(s)
	start := 0
	idx := 0

	for i := 0; i <= sIdx; i++ {
		_, size := utf8.DecodeRune(b[idx:])
		start = idx
		idx += size
	}

	return s[start:idx]
}
