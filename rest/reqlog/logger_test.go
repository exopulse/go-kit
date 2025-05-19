package reqlog

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestSetLogger(t *testing.T) {
	// create a gin context with a request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	c, _ := gin.CreateTestContext(w)

	c.Request = req

	// create a test logger with a unique field
	testValue := "test-value"

	var buf bytes.Buffer

	logger := zerolog.New(&buf).With().Str("test-field", testValue).Logger()

	// set the logger in the context
	SetLogger(c, logger)

	// retrieve the logger from the context
	retrievedLogger := RequestLogger(c)

	// verify the logger was set correctly by checking if the output contains our test field
	buf.Reset()
	retrievedLogger.Info().Str("another-field", "value").Msg("test message")

	require.Contains(t, buf.String(), testValue)
}

func TestRequestLogger(t *testing.T) {
	// test with logger in context
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	c, _ := gin.CreateTestContext(w)

	c.Request = req

	testValue := "test-value"

	var buf bytes.Buffer

	logger := zerolog.New(&buf).With().Str("test-field", testValue).Logger()

	SetLogger(c, logger)

	retrievedLogger := RequestLogger(c)

	require.NotNil(t, retrievedLogger)

	// verify it's the same logger
	buf.Reset()
	retrievedLogger.Info().Msg("test message")

	require.Contains(t, buf.String(), testValue)

	// test with no logger in context
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/", nil)
	c, _ = gin.CreateTestContext(w)

	c.Request = req

	retrievedLogger = RequestLogger(c)

	require.NotNil(t, retrievedLogger)
	require.True(t, retrievedLogger.GetLevel() == zerolog.Disabled)
}

func TestLogger(t *testing.T) {
	// test with logger in context
	testValue := "test-value"

	var buf bytes.Buffer

	logger := zerolog.New(&buf).With().Str("test-field", testValue).Logger()
	ctx := logger.WithContext(context.Background())

	retrievedLogger := Logger(ctx)

	require.NotNil(t, retrievedLogger)

	// verify it's the same logger
	buf.Reset()

	retrievedLogger.Info().Msg("test message")

	require.Contains(t, buf.String(), testValue)

	// test with no logger in context
	ctx = context.Background()
	retrievedLogger = Logger(ctx)

	require.NotNil(t, retrievedLogger)
	require.True(t, retrievedLogger.GetLevel() == zerolog.Disabled)
}

func TestUpdateLogger(t *testing.T) {
	// create a gin context with a request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	c, _ := gin.CreateTestContext(w)

	c.Request = req

	// set initial logger
	initialValue := "initial-value"

	var buf bytes.Buffer

	logger := zerolog.New(&buf).With().Str("initial-field", initialValue).Logger()

	SetLogger(c, logger)

	// update the logger
	updatedValue := "updated-value"

	UpdateLogger(c, func(current *zerolog.Logger) zerolog.Logger {
		return current.With().Str("updated-field", updatedValue).Logger()
	})

	// retrieve the updated logger
	retrievedLogger := RequestLogger(c)

	// verify the logger was updated correctly
	buf.Reset()
	retrievedLogger.Info().Msg("test message")

	require.Contains(t, buf.String(), initialValue)
	require.Contains(t, buf.String(), updatedValue)
}

func TestUpdateContextLogger(t *testing.T) {
	// create a context with logger
	initialValue := "initial-value"

	var buf bytes.Buffer

	logger := zerolog.New(&buf).With().Str("initial-field", initialValue).Logger()
	ctx := logger.WithContext(context.Background())

	// update the logger
	updatedValue := "updated-value"
	updatedCtx := updateContextLogger(ctx, func(current *zerolog.Logger) zerolog.Logger {
		return current.With().Str("updated-field", updatedValue).Logger()
	})

	// retrieve the updated logger
	retrievedLogger := Logger(updatedCtx)

	// verify the logger was updated correctly
	buf.Reset()
	retrievedLogger.Info().Msg("test message")

	require.Contains(t, buf.String(), initialValue)
	require.Contains(t, buf.String(), updatedValue)

	// test with no logger in context
	ctx = context.Background()

	var newBuf bytes.Buffer

	updatedCtx = updateContextLogger(ctx, func(current *zerolog.Logger) zerolog.Logger {
		return zerolog.New(&newBuf).With().Str("updated-field", updatedValue).Logger()
	})

	retrievedLogger = Logger(updatedCtx)

	require.NotNil(t, retrievedLogger)

	// verify the logger was updated correctly
	retrievedLogger.Info().Msg("test message")

	require.Contains(t, newBuf.String(), updatedValue)
}
