package env

import (
	"errors"
	"io/fs"
	"os"
	"path"
	"strings"
)

// resolveFilePath takes a file path as input and returns the absolute path.
// If the file starts with a /, it is considered an absolute path.
// Otherwise, it is considered a relative path and is resolved using the getwd function.
// Otherwise, the resolved path is returned.
func resolveFilePath(file string, getwd func() (string, error)) (string, error) {
	// assume the file is an absolute path
	if strings.HasPrefix(file, "/") {
		return file, nil
	}

	pwd, err := getwd()
	if err != nil {
		return "", err
	}

	return path.Join(pwd, file), nil
}

// readOptionalFile reads the content of the file at the given path.
// If the file does not exist, it returns an empty string.
func readOptionalFile(file string) (string, error) {
	content, err := os.ReadFile(path.Clean(file))
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return "", nil
		}

		//nolint:wrapcheck // return the error as is
		return "", err
	}

	return string(content), nil
}
