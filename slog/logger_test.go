package slog

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLevels(t *testing.T) {
	t.Parallel()

	require.NotPanics(t, func() {
		Info().Msg("info")
		Debug().Msg("debug")
		Warn().Msg("warn")
		Error().Msg("error")
	})
}

func TestContext(t *testing.T) {
	t.Parallel()

	t.Run("no-slog", func(t *testing.T) {
		ctx := context.Background()
		loggerFromCtx := FromContext(ctx)

		require.NotNil(t, loggerFromCtx)
	})

	t.Run("with-slog", func(t *testing.T) {
		ctx := WithContext(context.Background())
		loggerFromCtx := FromContext(ctx)

		require.NotNil(t, loggerFromCtx)
	})
}
