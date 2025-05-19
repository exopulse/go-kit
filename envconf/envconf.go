// Package envconf parses environment variables into a custom structs.
// Parser recognizes following boolean flags:
//   - true, false, on, yes, off, no
//
// Parser also recognizes following custom formats:
//   - timex.Duration
package envconf

import (
	"reflect"
	"strconv"

	envs "github.com/caarlos0/env/v7"
	"github.com/exopulse/go-kit/timex"
)

// Parse parses a struct containing `env` tags and loads its values from environment variables.
//
//nolint:wrapcheck // no need to wrap these errors
func Parse(v any) error {
	return envs.ParseWithFuncs(v, customParsers())
}

func customParsers() map[reflect.Type]envs.ParserFunc {
	return map[reflect.Type]envs.ParserFunc{
		reflect.TypeOf(true): func(v string) (any, error) {
			switch v {
			case "on", "yes":
				return true, nil
			case "off", "no":
				return false, nil
			default:
				return strconv.ParseBool(v)
			}
		},
		reflect.TypeOf(timex.Duration(0)): func(v string) (any, error) {
			return timex.ParseDuration(v)
		},
	}
}
