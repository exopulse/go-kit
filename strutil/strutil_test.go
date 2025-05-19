package strutil

import (
	"testing"

	. "github.com/stretchr/testify/require"
)

func TestUnquote(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"empty", "", ""},
		{"empty-dbl", `""`, ""},
		{"empty-sng", `''`, ""},
		{"value-dbl", `"value"`, "value"},
		{"value-sng", `'value'`, "value"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unquote(tt.arg); got != tt.want {
				t.Errorf("Unquote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOverride(t *testing.T) {
	Equal(t, "v1", Override("v1", "d1"))
	Equal(t, "d1", Override("", "d1"))
	Equal(t, " ", Override(" ", "d1"))
}

func TestLimitLength(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		maxLen int
		want   string
	}{
		{"ok", "text", 10, "text"},
		{"at-limit", "at-limit", 8, "at-limit"},
		{"over", "over-the-limit", 8, "over-the"},
		{"utf-split", "šđčćžŠĐČĆŽ", 5, "šđčćž"},
		{"utf-mix-split", "asdŽŠČ", 4, "asdŽ"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LimitLength(tt.s, tt.maxLen); got != tt.want {
				t.Errorf("LimitLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Uncapitalize(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input string
		want  string
	}{
		"empty": {
			input: "",
			want:  "",
		},
		"lower": {
			input: "lower",
			want:  "lower",
		},
		"Upper": {
			input: "Upper",
			want:  "upper",
		},
		"UPPER": {
			input: "UPPER",
			want:  "uPPER",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := Uncapitalize(tt.input)

			Equal(t, tt.want, got)
		})
	}
}
