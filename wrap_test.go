package flagvar_test

import (
	"flag"
	"reflect"
	"testing"

	"github.com/sgreben/flagvar"
)

func TestWrapPointer(t *testing.T) {
	sv := flagvar.Strings{}
	var p flag.Value = &sv
	fv := flagvar.WrapPointer{
		Value: &p,
	}
	var fs flag.FlagSet
	fs.Var(&fv, "wrap-pointer", "")

	err := fs.Parse([]string{"-wrap-pointer", "abc"})
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(sv.Values, []string{"abc"}) {
		t.Fail()
	}
}

func TestWrapFunc(t *testing.T) {
	sv := &flagvar.Strings{}
	fv := flagvar.WrapFunc(func() flag.Value { return sv })
	var fs flag.FlagSet
	fs.Var(&fv, "wrap-func", "")

	err := fs.Parse([]string{"-wrap-func", "abc"})
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(sv.Values, []string{"abc"}) {
		t.Fail()
	}
}

func TestWrap(t *testing.T) {
	updated := 0
	sv := &flagvar.Strings{}
	fv := flagvar.Wrap{
		Value: sv,
		Updated: func() {
			updated++
		},
	}
	var fs flag.FlagSet
	fs.Var(&fv, "wrap", "")

	err := fs.Parse([]string{"-wrap", "abc", "-wrap", "def"})
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(sv.Values, []string{"abc", "def"}) {
		t.Fail()
	}
	if updated != 2 {
		t.Fail()
	}
}
