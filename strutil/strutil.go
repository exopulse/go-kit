package strutil

import (
	"strings"
	"unicode/utf8"
)

const minQuotedExpressionLen = 2

// Unquote removes single or double quotes around input string.
func Unquote(s string) string {
	if len(s) >= minQuotedExpressionLen {
		if q := s[0]; s[len(s)-1] == q && (q == '"' || q == '\'') {
			return s[1 : len(s)-1]
		}
	}

	return s
}

// Override returns value if not empty. Otherwise, it returns defaultValue.
// Value containing only spaces is not considered to be empty.
func Override(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}

	return value
}

// LimitLength limits the string length to a specified value.
func LimitLength(s string, maxLen int) string {
	if utf8.RuneCountInString(s) > maxLen {
		rs := []rune(s)

		return string(rs[:maxLen])
	}

	return s
}

// Uncapitalize returns a copy of the string s with the first
// Unicode letter mapped to its lower case.
func Uncapitalize(s string) string {
	if s == "" {
		return s
	}

	return strings.ToLower(string(s[0])) + s[1:]
}
