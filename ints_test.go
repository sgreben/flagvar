package flagvar_test

import (
	"flag"
	"reflect"
	"testing"

	"github.com/sgreben/flagvar"
)

func TestInts(t *testing.T) {
	fv := flagvar.Ints{}
	var fs flag.FlagSet
	fs.Var(&fv, "ints", "")

	err := fs.Parse([]string{"-ints", "123", "-ints", "9090"})
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(fv.Values, []int64{123, 9090}) {
		t.Fail()
	}
}

func TestIntsBitSize(t *testing.T) {
	fv := flagvar.Ints{BitSize: 32}
	var fs flag.FlagSet
	fs.Var(&fv, "ints", "")

	err := fs.Parse([]string{"-ints", "123", "-ints", "9090"})
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(fv.Values, []int64{123, 9090}) {
		t.Fail()
	}
}

func TestIntsBase(t *testing.T) {
	fv := flagvar.Ints{BitSize: 32, Base: 16}
	var fs flag.FlagSet
	fs.Var(&fv, "ints", "")

	err := fs.Parse([]string{"-ints", "F", "-ints", "0"})
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(fv.Values, []int64{15, 0}) {
		t.Fail()
	}
}

func TestIntsFail(t *testing.T) {
	fv := flagvar.Ints{}
	var fs flag.FlagSet
	fs.Var(&fv, "ints", "")

	err := fs.Parse([]string{"-ints", "abc"})
	if err == nil {
		t.Fail()
	}
}
