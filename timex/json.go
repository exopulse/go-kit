package timex

import (
	"bytes"
)

// MarshalJSON converts this value to JSON value.
// Non-empty value is double-quoted.
// nosemgrep // disable semgrep-go.marshal-json-pointer-receiver.
func (d *Duration) MarshalJSON() ([]byte, error) {
	w := new(bytes.Buffer)

	_ = w.WriteByte('"')
	_, _ = w.WriteString(d.String())
	_ = w.WriteByte('"')

	return w.Bytes(), nil
}

// UnmarshalJSON converts JSON value to Duration.
// Non-empty value is unquoted and converted to Duration.
// Empty value is converted to idn.Nil.
func (d *Duration) UnmarshalJSON(b []byte) error {
	s := string(b)
	l := len(s)

	if l >= 2 && b[0] == '"' && b[l-1] == '"' {
		s = s[1 : l-1]
	}

	pd, err := ParseDuration(s)
	if err != nil {
		return err
	}

	*d = pd

	return err
}
