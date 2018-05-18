package flagvar_test

import (
	"flag"
	"reflect"
	"testing"

	"github.com/sgreben/flagvar"
)

func TestEnum(t *testing.T) {
	fv := flagvar.Enum{Choices: []string{"first", "second"}}
	var fs flag.FlagSet
	fs.Var(&fv, "enum", "")

	err := fs.Parse([]string{"-enum", "first", "-enum", "second"})
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(fv.Value, "second") {
		t.Fail()
	}
}

func TestEnumFail(t *testing.T) {
	fv := flagvar.Enum{Choices: []string{"first", "second"}}
	var fs flag.FlagSet
	fs.Var(&fv, "enum", "")

	err := fs.Parse([]string{"-enum", "third"})
	if err == nil {
		t.Fail()
	}
}

func TestEnums(t *testing.T) {
	fv := flagvar.Enums{Choices: []string{"first", "second"}}
	var fs flag.FlagSet
	fs.Var(&fv, "enum", "")

	err := fs.Parse([]string{"-enum", "first"})
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(fv.Values, []string{"first"}) {
		t.Fail()
	}
}

func TestEnumsFail(t *testing.T) {

	fv := flagvar.Enums{Choices: []string{"first", "second"}}
	var fs flag.FlagSet
	fs.Var(&fv, "enum", "")

	err := fs.Parse([]string{"-enum", "third"})
	if err == nil {
		t.Fail()
	}
}
