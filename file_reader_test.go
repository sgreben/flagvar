package flagvar_test

import (
	"flag"
	"testing"

	"github.com/sgreben/flagvar"
)

func TestFileReader(t *testing.T) {
	fr := flagvar.FileReader{}
	var fs flag.FlagSet
	fs.Var(&fr, "f", "input file")

	err := fs.Parse([]string{"-f", "./noSuchFile.tpl"})
	if err == nil {
		t.Fail()
	}
}
