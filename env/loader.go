package env

import (
	"fmt"
	"os"
	"strings"
)

// Loader loads environment variables from a file.
type Loader struct {
	setter     func(string, string) error
	getwd      func() (string, error)
	fileReader func(path string) (string, error)
}

// NewLoader creates a new Loader.
func NewLoader() *Loader {
	return &Loader{
		setter:     os.Setenv,
		getwd:      os.Getwd,
		fileReader: readOptionalFile,
	}
}

// LoadOptional loads the environment variables from the given file.
// It ignores the file if it does not exist.
func (l *Loader) LoadOptional(file string) error {
	file, err := resolveFilePath(file, l.getwd)
	if err != nil {
		return fmt.Errorf("failed to resolve file path: %w", err)
	}

	content, err := l.fileReader(file)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	return apply(content, l.setter)
}

// Apply loads the environment variables from the given content.
func (l *Loader) Apply(content string) error {
	return apply(content, l.setter)
}

// apply sets the environment variables from the given content.
func apply(content string, setter func(key, value string) error) error {
	for _, line := range strings.Split(content, "\n") {
		key, value := parseLine(line)
		if key == "" {
			continue
		}

		if err := setter(key, value); err != nil {
			return err
		}
	}

	return nil
}
