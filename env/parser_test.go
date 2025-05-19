package env

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_parseLine(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input     string
		wantKey   string
		wantValue string
	}{
		"empty": {
			input:     "",
			wantKey:   "",
			wantValue: "",
		},
		"no-separator": {
			input:     "hello",
			wantKey:   "",
			wantValue: "",
		},
		"no-value": {
			input:     "hello=",
			wantKey:   "hello",
			wantValue: "",
		},
		"no-key": {
			input:     "=world",
			wantKey:   "",
			wantValue: "",
		},
		"key-value": {
			input:     "hello=world",
			wantKey:   "hello",
			wantValue: "world",
		},
		"key-value-spaces": {
			input:     " hello = world ",
			wantKey:   "hello",
			wantValue: "world",
		},
		"key-value-quoted": {
			input:     "hello='world'",
			wantKey:   "hello",
			wantValue: "world",
		},
		"key-value-quoted-spaces": {
			input:     " hello = 'world' ",
			wantKey:   "hello",
			wantValue: "world",
		},
		"multi-separator": {
			input:     "hello=world=again",
			wantKey:   "hello",
			wantValue: "world=again",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			key, value := parseLine(tt.input)

			require.Equal(t, tt.wantKey, key)
			require.Equal(t, tt.wantValue, value)
		})
	}
}

func Test_trimComment(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input string
		want  string
	}{
		"empty": {
			input: "",
			want:  "",
		},
		"no-comment": {
			input: "hello",
			want:  "hello",
		},
		"comment": {
			input: "hello # world",
			want:  "hello ",
		},
		"comment-only": {
			input: "# hello",
			want:  "",
		},
		"comment-spaces": {
			input: " # hello",
			want:  " ",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tt.want, trimComment(tt.input))
		})
	}
}

func Test_unquote(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input string
		want  string
	}{
		"empty": {
			input: "",
			want:  "",
		},
		"single-quote": {
			input: "'hello'",
			want:  "hello",
		},
		"double-quote": {
			input: "\"hello\"",
			want:  "hello",
		},
		"no-quote": {
			input: "hello",
			want:  "hello",
		},
		"no-quote-spaces": {
			input: "hello world",
			want:  "hello world",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tt.want, unquote(tt.input))
		})
	}
}
