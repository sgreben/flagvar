package flagvar

import (
	"fmt"
	"text/template"
)

type Template struct {
	Root *template.Template

	Value *template.Template
	Text  string
}

// Set is flag.Value.Set
func (fv *Template) Set(v string) error {
	root := fv.Root
	if root == nil {
		root = template.New("")
	}
	t, err := root.New(fmt.Sprint(fv)).Parse(v)
	if err == nil {
		fv.Value = t
	}
	return err
}

func (fv *Template) String() string {
	return fv.Text
}

type Templates struct {
	Root *template.Template

	Values []*template.Template
	Texts  []string
}

// Set is flag.Value.Set
func (fv *Templates) Set(v string) error {
	root := fv.Root
	if root == nil {
		root = template.New("")
	}
	t, err := root.New(fmt.Sprint(fv)).Parse(v)
	if err == nil {
		fv.Texts = append(fv.Texts, v)
		fv.Values = append(fv.Values, t)
	}
	return err
}

func (fv *Templates) String() string {
	return fmt.Sprint(fv.Texts)
}