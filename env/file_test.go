package env

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_resolveFilePath(t *testing.T) {
	t.Parallel()

	t.Run("absolute-path", func(t *testing.T) {
		t.Parallel()

		file := "/tmp/.env"

		got, err := resolveFilePath(file, func() (string, error) {
			return "", nil
		})

		require.NoError(t, err)
		require.Equal(t, file, got)
	})

	t.Run("relative-path", func(t *testing.T) {
		t.Parallel()

		file := ".env"
		pwd := "/tmp"

		got, err := resolveFilePath(file, func() (string, error) {
			return pwd, nil
		})

		require.NoError(t, err)
		require.Equal(t, "/tmp/.env", got)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		forcedError := errors.New("forced-error")

		_, err := resolveFilePath(".env", func() (string, error) {
			return "", forcedError
		})

		require.ErrorIs(t, err, forcedError)
	})
}

func Test_readOptionalFile(t *testing.T) {
	t.Parallel()

	t.Run("file-exists", func(t *testing.T) {
		t.Parallel()

		file := "testdata/.env"

		got, err := readOptionalFile(file)

		require.NoError(t, err)
		require.Equal(t, "KEY1=value1\n", got)
	})

	t.Run("file-not-exists", func(t *testing.T) {
		t.Parallel()

		file := "testdata/unknown"

		got, err := readOptionalFile(file)

		require.NoError(t, err)
		require.Empty(t, got)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		// read directory as file, which should fail
		_, err := readOptionalFile("testdata")

		require.Error(t, err)
	})
}
