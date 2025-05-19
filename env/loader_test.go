package env

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoader_LoadOptional(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		envs := map[string]string{}

		envLoader := NewLoader()

		envLoader.setter = func(key, value string) error {
			envs[key] = value

			return nil
		}

		err := envLoader.LoadOptional("testdata/.env")

		require.NoError(t, err)
		require.Equal(t, map[string]string{"KEY1": "value1"}, envs)
	})

	t.Run("path-resolve-error", func(t *testing.T) {
		t.Parallel()

		envLoader := NewLoader()

		envLoader.getwd = func() (string, error) {
			return "", errors.New("forced-error")
		}

		err := envLoader.LoadOptional(".env")

		require.Error(t, err)
	})

	t.Run("file-loading-error", func(t *testing.T) {
		t.Parallel()

		envLoader := NewLoader()

		envLoader.fileReader = func(path string) (string, error) {
			return "", errors.New("forced-error")
		}

		err := envLoader.LoadOptional("testdata/.env")

		require.Error(t, err)
	})
}

func TestLoader_Apply(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		envs := map[string]string{}

		envLoader := NewLoader()

		envLoader.setter = func(key, value string) error {
			envs[key] = value

			return nil
		}

		err := envLoader.Apply("KEY1=value1")

		require.NoError(t, err)
		require.Equal(t, map[string]string{"KEY1": "value1"}, envs)
	})

	t.Run("setter-failure", func(t *testing.T) {
		t.Parallel()

		envLoader := NewLoader()

		envLoader.setter = func(key, value string) error {
			return errors.New("forced-error")
		}

		err := envLoader.Apply("KEY1=value1")

		require.Error(t, err)
	})
}

func Test_apply(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		setterCalled := false

		err := apply("", func(key, value string) error {
			setterCalled = true

			return nil
		})

		require.NoError(t, err)
		require.False(t, setterCalled)
	})

	t.Run("key-values-applied", func(t *testing.T) {
		t.Parallel()

		env := map[string]string{}

		err := apply("key=value\n\nkey2=value2", func(key, value string) error {
			env[key] = value

			return nil
		})

		require.NoError(t, err)
		require.Equal(t, map[string]string{"key": "value", "key2": "value2"}, env)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		forcedError := errors.New("forced-error")

		err := apply("key=value", func(key, value string) error {
			return forcedError
		})

		require.ErrorIs(t, err, forcedError)
	})
}
