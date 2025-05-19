package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

// mockRoute implements the Route interface for testing.
type mockRoute struct {
	registerRoutesCalled bool
	handler              gin.HandlerFunc
}

func newMockRoute(handler gin.HandlerFunc) *mockRoute {
	return &mockRoute{
		handler: handler,
	}
}

func (m *mockRoute) RegisterRoutes(group *gin.RouterGroup) {
	m.registerRoutesCalled = true

	group.GET("/test", m.handler)
}

func TestNew(t *testing.T) {
	// test without middleware
	r := New()

	require.NotNil(t, r)
	require.NotNil(t, r.rtr)

	// test with middleware
	middlewareCalled := false
	middleware := func(c *gin.Context) {
		middlewareCalled = true

		c.Next()
	}

	r = New(middleware)

	require.NotNil(t, r)
	require.NotNil(t, r.rtr)

	// verify middleware is called
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	r.ServeHTTP(w, req)

	require.True(t, middlewareCalled)
}

func TestRouter_ServeHTTP(t *testing.T) {
	r := New()

	// create a mock route that responds with a specific status code
	responseStatus := http.StatusOK
	responseBody := "test response"
	mockHandler := func(c *gin.Context) {
		c.String(responseStatus, responseBody)
	}

	mockR := newMockRoute(mockHandler)

	r.RegisterRoutes("/api", mockR)

	// test the ServeHTTP method
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/test", nil)

	r.ServeHTTP(w, req)

	require.Equal(t, responseStatus, w.Code)
	require.Equal(t, responseBody, w.Body.String())
}

func TestRouter_RegisterRoutes(t *testing.T) {
	r := New()

	// create a mock route
	mockHandler := func(c *gin.Context) {
		c.Status(http.StatusOK)
	}

	mockR := newMockRoute(mockHandler)

	// register the mock route
	r.RegisterRoutes("/api", mockR)

	// verify the route was registered
	require.True(t, mockR.registerRoutesCalled)

	// test that the route works
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/test", nil)

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}

func TestRouter_GetRoutes(t *testing.T) {
	r := New()

	// register a few routes
	mockHandler := func(c *gin.Context) {}

	mockR1 := newMockRoute(mockHandler)

	r.RegisterRoutes("/api/v1", mockR1)

	mockR2 := &mockRoute{handler: mockHandler}

	r.RegisterRoutes("/api/v2", mockR2)

	routes := r.GetRoutes()

	// verify we have at least the two routes we registered
	require.GreaterOrEqual(t, len(routes), 2)

	// check that our routes are included
	foundV1 := false
	foundV2 := false

	for _, route := range routes {
		if route.Method == "GET" && route.Path == "/api/v1/test" {
			foundV1 = true
		}

		if route.Method == "GET" && route.Path == "/api/v2/test" {
			foundV2 = true
		}
	}

	require.True(t, foundV1, "Expected to find /api/v1/test route")
	require.True(t, foundV2, "Expected to find /api/v2/test route")
}

func TestRouteInfo(t *testing.T) {
	info := RouteInfo{
		Method: "GET",
		Path:   "/test",
	}

	require.Equal(t, "GET", info.Method)
	require.Equal(t, "/test", info.Path)
}
