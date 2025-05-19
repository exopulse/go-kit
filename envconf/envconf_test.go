package envconf

import (
	"os"
	"testing"

	"github.com/exopulse/go-kit/timex"
	. "github.com/stretchr/testify/require"
)

type conf struct {
	Name     string         `env:"NAME"`
	Int      int            `env:"INT"`
	True     bool           `env:"TRUE"`
	False    bool           `env:"FALSE"`
	On       bool           `env:"ON"`
	Off      bool           `env:"OFF"`
	Yes      bool           `env:"YES"`
	No       bool           `env:"NO"`
	Duration timex.Duration `env:"DURATION"`
}

func TestParse(t *testing.T) {
	setEnv("NAME", "Name")
	setEnv("INT", "1234")
	setEnv("TRUE", "true")
	setEnv("FALSE", "false")
	setEnv("ON", "on")
	setEnv("OFF", "off")
	setEnv("YES", "yes")
	setEnv("NO", "no")
	setEnv("DURATION", "2y")

	cf := conf{}

	err := Parse(&cf)
	NoError(t, err)

	Equal(t, "Name", cf.Name)
	Equal(t, 1234, cf.Int)
	True(t, cf.True)
	False(t, cf.False)
	True(t, cf.On)
	False(t, cf.Off)
	True(t, cf.Yes)
	False(t, cf.No)
	EqualValues(t, parseDuration("2y"), cf.Duration)
}

func TestParse_Error(t *testing.T) {
	setEnv("INT", "asd")

	cf := conf{}

	err := Parse(&cf)
	Error(t, err)
}

func TestParse_Boolean_InvalidCase(t *testing.T) {
	setEnv("YES", "Yes")

	cf := conf{}

	err := Parse(&cf)
	Error(t, err)
}

func setEnv(name, value string) {
	if err := os.Setenv(name, value); err != nil {
		panic(err)
	}
}

func parseDuration(s string) timex.Duration {
	duration, err := timex.ParseDuration(s)
	if err != nil {
		panic(err)
	}

	return duration
}
