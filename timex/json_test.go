package timex

import (
	"encoding/json"
	"testing"

	. "github.com/stretchr/testify/require"
)

type object1 struct {
	Duration Duration `json:"duration"`
}

type object2 struct {
	Duration *Duration `json:"duration"`
}

func TestJSON_Encode(t *testing.T) {
	tests := []struct {
		object any
		want   string
	}{
		{&object1{mustParse("3mo")}, `{"duration":"3mo"}`},
		{&object1{}, `{"duration":"0s"}`},
		{&object2{mustParseRef("3mo")}, `{"duration":"3mo"}`},
		{&object2{}, `{"duration":null}`},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			bytes, err := json.Marshal(&tt.object)
			if err != nil {
				panic(err)
			}

			Equal(t, tt.want, string(bytes))
		})
	}
}

func TestJSON_Decode(t *testing.T) {
	tests := []struct {
		input    string
		receiver any
		want     any
		wantErr  bool
	}{
		{`{"duration":"3mo"}`, &object1{}, &object1{mustParse("3mo")}, false},
		{`{"duration":"0s"}`, &object1{}, &object1{}, false},
		{`{"duration":"3mo"}`, &object2{}, &object2{mustParseRef("3mo")}, false},
		{`{"duration":null}`, &object2{}, &object2{}, false},
		{`{"duration":"error"}`, &object2{}, nil, true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			err := json.Unmarshal([]byte(tt.input), &tt.receiver)

			Equal(t, tt.wantErr, err != nil)

			if err != nil {
				return
			}

			Equal(t, tt.want, tt.receiver)
		})
	}
}

func mustParse(s string) Duration {
	d, err := ParseDuration(s)
	if err != nil {
		panic(err)
	}

	return d
}

func mustParseRef(s string) *Duration {
	d := mustParse(s)

	return &d
}
