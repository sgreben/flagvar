package flagvar

import (
	"path/filepath"
	"strings"

	"github.com/gobwas/glob"
)

// Glob is a `flag.Value` for glob syntax arguments.
type Glob struct {
	Value glob.Glob
	Text  string
}

// Set is flag.Value.Set
func (fv *Glob) Set(v string) error {
	g, err := glob.Compile(v, filepath.Separator)
	if err != nil {
		return err
	}
	fv.Text = v
	fv.Value = g
	return nil
}

func (fv *Glob) String() string {
	return fv.Text
}

// Globs is a `flag.Value` for glob syntax arguments.
type Globs struct {
	Values []glob.Glob
	Texts  []string
}

// Set is flag.Value.Set
func (fv *Globs) Set(v string) error {
	g, err := glob.Compile(v, filepath.Separator)
	if err != nil {
		return err
	}
	fv.Texts = append(fv.Texts, v)
	fv.Values = append(fv.Values, g)
	return nil
}

func (fv *Globs) String() string {
	return strings.Join(fv.Texts, ",")
}
