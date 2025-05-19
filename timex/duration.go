package timex

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Duration shadows built-in time.Duration.
//
//nolint:recvcheck // driver.Valuer requires a value receiver.
type Duration time.Duration

const (
	day   = int64(time.Hour) * 24
	month = day * 30
	year  = day * 365
)

// ParseDuration parses a duration string.
// Time units "d", "mo" and "y" are treated specifically. In that case only integer values are accepted.
// Other units are passed through to time.ParseDuration.
// As a special cases, units are treated as:
// - day (d) - 24 hours
// - month (mo) - 30 days
// - year (y) - 365 days
// Negative duration is supported.
//
//nolint:cyclop // I can live with this function the way it is
func ParseDuration(s string) (Duration, error) {
	unitAt := strings.IndexAny(s, "dmy")

	if unitAt == -1 || (s[unitAt:unitAt+1] == "m" && s[unitAt:] != "mo") {
		return parseAsBuiltInDuration(s)
	}

	orig := s
	neg := false

	var d int64

	if s != "" {
		c := s[0]

		if c == '-' || c == '+' {
			neg = c == '-'
			s = s[1:]
			unitAt--
		}
	}

	n, err := strconv.Atoi(s[:unitAt])
	if err != nil {
		return 0, errors.Wrapf(err, "invalid duration string '%s'", orig)
	}

	d = int64(n)
	unit := s[unitAt:]

	switch unit {
	case "d":
		d *= day
	case "mo":
		d *= month
	case "y":
		d *= year
	default:
		return 0, errors.Errorf("invalid duration value '%s' in '%s'", unit, orig)
	}

	if neg {
		d = -d
	}

	return Duration(d), nil
}

func parseAsBuiltInDuration(s string) (Duration, error) {
	d, err := time.ParseDuration(s)

	return Duration(d), errors.Wrap(err, "failed to parse duration")
}

// Duration returns time.Duration.
func (d Duration) Duration() time.Duration {
	return time.Duration(d)
}

func (d Duration) String() string {
	v := int64(d)
	vp := v

	if vp < 0 {
		vp = -vp
	}

	// we will format only integer values, because twe could only parse integers
	// anything other than that came in as a time.Duration
	switch {
	case vp >= year && vp%year == 0:
		return fmt.Sprintf("%dy", v/year)
	case vp >= month && vp%month == 0:
		return fmt.Sprintf("%dmo", v/month)
	case vp >= day && vp%day == 0:
		return fmt.Sprintf("%dd", v/day)
	}

	return time.Duration(d).String()
}
