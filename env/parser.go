package env

import "strings"

// parseLine takes a line of string as input and returns a key-value pair.
// It handles comments, separators, and quotations in the line.
// If the line cannot be parsed into a key-value pair, it returns two empty strings.
func parseLine(line string) (string, string) {
	const separator = "="

	line = trimComment(line)

	key, value, found := strings.Cut(line, separator)
	if !found {
		return "", ""
	}

	key = strings.TrimSpace(key)
	if key == "" {
		return "", ""
	}

	value = unquote(strings.TrimSpace(value))

	return key, value
}

// trimComment strips a trailing "# ..." comment from the line.
// A '#' inside a single- or double-quoted span is not treated as a comment,
// so quoted values may contain '#' (e.g. hello="pass#word").
func trimComment(line string) string {
	const comment = '#'

	var quote byte

	for i := range len(line) {
		c := line[i]

		switch {
		case quote != 0:
			if c == quote {
				quote = 0
			}
		case c == '"' || c == '\'':
			quote = c
		case c == comment:
			return line[:i]
		}
	}

	return line
}

func unquote(s string) string {
	const minLength = 2

	if len(s) < minLength {
		return s
	}

	if s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}

	if s[0] == '\'' && s[len(s)-1] == '\'' {
		return s[1 : len(s)-1]
	}

	return s
}
