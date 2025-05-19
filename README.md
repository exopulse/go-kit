# ExoPulse Go-Kit

A comprehensive toolkit of Go utilities and helpers designed to simplify common tasks in Go applications.

## Overview

ExoPulse Go-Kit is a collection of lightweight, reusable Go packages that provide solutions for common programming tasks. The toolkit is designed with simplicity, performance, and modern Go practices in mind.

## Features

- **Environment Configuration** (`envconf`): Parse environment variables into Go structs with extended boolean flag support and custom formats.
- **Host Utilities** (`hostutil`): Tools for working with host-related functionality.
- **HTTP Server** (`httpd`): Utilities for HTTP server implementation.
- **REST Helpers** (`rest`):
  - `reqlog`: Request logging middleware and utilities for Gin framework
  - `router`: Simplified router implementation for Gin-based applications
- **Structured Logging** (`slog`): Zerolog-based structured logging with context support.
- **String Utilities** (`strutil`): Helper functions for string manipulation.
- **Time Extensions** (`timex`): Extended time functionality and duration parsing.

## Installation

```bash
go get github.com/exopulse/go-kit
```

## Usage Examples

### Environment Configuration

```go
package main

import (
    "fmt"
    "github.com/exopulse/go-kit/envconf"
    "github.com/exopulse/go-kit/timex"
)

type Config struct {
    Debug    bool          `env:"DEBUG"`
    Timeout  timex.Duration `env:"TIMEOUT"`
    LogLevel string        `env:"LOG_LEVEL" envDefault:"info"`
}

func main() {
    var cfg Config
    if err := envconf.Parse(&cfg); err != nil {
        panic(err)
    }
    
    fmt.Printf("Config: %+v\n", cfg)
}
```

### Request Logging

```go
package main

import (
    "github.com/exopulse/go-kit/rest/reqlog"
    "github.com/exopulse/go-kit/slog"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.New()
    
    // Add request logger middleware
    r.Use(func(c *gin.Context) {
        logger := slog.Global
        reqlog.SetLogger(c, logger)
        c.Next()
    })
    
    r.GET("/example", func(c *gin.Context) {
        // Get logger from context
        logger := reqlog.RequestLogger(c)
        logger.Info().Msg("Processing request")
        
        c.JSON(200, gin.H{"status": "ok"})
    })
    
    r.Run(":8080")
}
```

## Requirements

- Go 1.24 or higher

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
