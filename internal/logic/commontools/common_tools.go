package commontools

import (
	"unicode"
	"unicode/utf8"
)

func ToLowerFirstChar(s string) string {
	if s == "" {
		return s
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[size:]
}
