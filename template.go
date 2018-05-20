package flagvar

import (
	"fmt"
	"text/template"
)

// Template is a `flag.Value` for `text.Template` arguments.
// The value of the `Root` field is used as a root template when specified.
type Template struct {
	Root *template.Template

	Value *template.Template
	Text  string
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *Template) Help() string {
	return "a go template"
}

// Set is flag.Value.Set
func (fv *Template) Set(v string) error {
	root := fv.Root
	if root == nil {
		root = template.New("")
	}
	t, err := root.New(fmt.Sprintf("%T(%p)", fv, fv)).Parse(v)
	if err == nil {
		fv.Value = t
	}
	return err
}

func (fv *Template) String() string {
	return fv.Text
}

// Templates is a `flag.Value` for `text.Template` arguments.
// The value of the `Root` field is used as a root template when specified.
type Templates struct {
	Root *template.Template

	Values []*template.Template
	Texts  []string
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *Templates) Help() string {
	return "a go template"
}

// Set is flag.Value.Set
func (fv *Templates) Set(v string) error {
	root := fv.Root
	if root == nil {
		root = template.New("")
	}
	t, err := root.New(fmt.Sprintf("%T(%p)", fv, fv)).Parse(v)
	if err == nil {
		fv.Texts = append(fv.Texts, v)
		fv.Values = append(fv.Values, t)
	}
	return err
}

func (fv *Templates) String() string {
	return fmt.Sprint(fv.Texts)
}
