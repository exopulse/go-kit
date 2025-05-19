package env

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAutoLoad(t *testing.T) {
	t.Parallel()

	require.NoError(t, AutoLoad(""))
}

func Test_autoLoad(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		envs := map[string]string{}

		envLoader := NewLoader()

		envLoader.setter = func(key, value string) error {
			envs[key] = value

			return nil
		}

		err := autoLoad("KEY2=value2", "", envLoader)

		require.NoError(t, err)
		require.Equal(t, map[string]string{"KEY2": "value2"}, envs)
	})

	t.Run("setenv-error", func(t *testing.T) {
		t.Parallel()

		envLoader := NewLoader()

		forcedError := errors.New("forced-error")

		envLoader.setter = func(key, value string) error {
			return forcedError
		}

		err := autoLoad("KEY2=value2", "", envLoader)

		require.ErrorIs(t, err, forcedError)
	})

	t.Run("file-error", func(t *testing.T) {
		t.Parallel()

		envLoader := NewLoader()

		forcedError := errors.New("forced-error")

		envLoader.fileReader = func(path string) (string, error) {
			return "", forcedError
		}

		err := autoLoad("", "", envLoader)

		require.ErrorIs(t, err, forcedError)
	})
}

func Test_resolveSelector(t *testing.T) {
	t.Parallel()

	t.Run("specified", func(t *testing.T) {
		t.Parallel()

		selector := "prod"

		require.Equal(t, selector, resolveSelector(selector))
	})

	t.Run("default", func(t *testing.T) {
		t.Parallel()

		require.Equal(t, "dev", resolveSelector(""))
	})
}
