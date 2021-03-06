package flagvar_test

import (
	"flag"
	"testing"
	"time"

	"github.com/sgreben/flagvar"
)

func TestTimeFormatISO8601(t *testing.T) {
	fv := flagvar.TimeFormat{}
	var fs flag.FlagSet
	fs.Var(&fv, "time-format", "")

	err := fs.Parse([]string{"-time-format", "RFC3339"})
	if err != nil {
		t.Error(err)
	}
	referenceTime := time.Unix(1594467527, 0)
	expected := referenceTime.Format(time.RFC3339)
	actual := referenceTime.Format(fv.Value)
	if actual != expected {
		t.Errorf("actual %q != %q (expected)", actual, expected)
	}
}

func TestTimeFormatFail(t *testing.T) {
	fv := flagvar.TimeFormat{}
	var fs flag.FlagSet
	fs.Var(&fv, "time-format", "")

	err := fs.Parse([]string{"-time-format", "[a-"})
	if err == nil {
		t.Fail()
	}
	t.Log(fv.Help())
}
