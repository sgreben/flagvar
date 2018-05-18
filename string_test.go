package flagvar_test

import (
	"flag"
	"reflect"
	"testing"

	"github.com/sgreben/flagvar"
)

func TestStrings(t *testing.T) {
	fv := flagvar.Strings{}
	var fs flag.FlagSet
	fs.Var(&fv, "strings", "")

	err := fs.Parse([]string{"-strings", "123", "-strings", "abc"})
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(fv.Values, []string{"123", "abc"}) {
		t.Fail()
	}
}

func TestStringSet(t *testing.T) {
	fv := flagvar.StringSet{}
	var fs flag.FlagSet
	fs.Var(&fv, "stringset", "")

	err := fs.Parse([]string{"-stringset", "123", "-stringset", "abc", "-stringset", "123"})
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(fv.Values(), []string{"123", "abc"}) {
		t.Fail()
	}
}
