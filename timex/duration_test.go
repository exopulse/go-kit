package timex

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParseDuration(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    Duration
		wantErr bool
	}{
		{"empty", "", 0, true},
		{"invalid", "x", 0, true},
		{"defaults-to-core", "2h", Duration(2 * time.Hour), false},
		{"one-day", "1d", Duration(time.Hour * 24), false},
		{"minus-one-day", "-1d", Duration(-time.Hour * 24), false},
		{"plus-one-day", "+1d", Duration(time.Hour * 24), false},
		{"one-month", "1mo", Duration(time.Hour * 24 * 30), false},
		{"minus-one-month", "-1mo", Duration(-time.Hour * 24 * 30), false},
		{"plus-one-month", "+1mo", Duration(time.Hour * 24 * 30), false},
		{"one-year", "1y", Duration(time.Hour * 24 * 365), false},
		{"minus-one-year", "-1y", Duration(-time.Hour * 24 * 365), false},
		{"plus-one-year", "+1y", Duration(time.Hour * 24 * 365), false},
		{"invalid-value", "1day", 0, true},
		{"invalid-minus-value", "-1day", 0, true},
		{"invalid-plus-value", "+1day", 0, true},
		{"missing-value", "day", 0, true},
		{"missing-minus-value", "-day", 0, true},
		{"missing-plus-value", "+day", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDuration(tt.s)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("ParseDuration() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDuration_String(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"days", "5d", "5d"},
		{"minus-days", "-5d", "-5d"},
		{"months", "5mo", "5mo"},
		{"minus-months", "-5mo", "-5mo"},
		{"years", "2y", "2y"},
		{"minus-years", "-2y", "-2y"},
		{"defaults-to-core", "2h2m", "2h2m0s"},
		{"26-hours", "28h", "28h0m0s"},
		{"96-hours", "96h", "4d"},
		{"98-hours", "98h", "98h0m0s"},
		{"60-days", "60d", "2mo"},
		{"65-days", "65d", "65d"},
		{"730-days", "730d", "2y"},
		{"731-days", "731d", "731d"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := ParseDuration(tt.s)
			if err != nil {
				panic(err)
			}

			if got := d.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDuration_Duration(t *testing.T) {
	d, err := ParseDuration("48d")
	if err != nil {
		panic(err)
	}

	require.Equal(t, 48*24*time.Hour, d.Duration())
}
