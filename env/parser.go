package env

import "strings"

// parseLine takes a line of string as input and returns a key-value pair.
// It handles comments, separators, and quotations in the line.
// If the line cannot be parsed into a key-value pair, it returns two empty strings.
func parseLine(line string) (string, string) {
	const separator = "="

	line = trimComment(line)

	separatorAt := strings.Index(line, separator)
	if separatorAt < 0 {
		return "", ""
	}

	key := strings.TrimSpace(line[:separatorAt])
	if key == "" {
		return "", ""
	}

	value := unquote(strings.TrimSpace(line[separatorAt+len(separator):]))

	return key, value
}

func trimComment(line string) string {
	const comment = "#"

	if commentAt := strings.Index(line, comment); commentAt != -1 {
		line = line[:commentAt]
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
