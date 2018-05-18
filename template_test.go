package flagvar_test

import (
	"flag"
	"reflect"
	"testing"
	"text/template"

	"github.com/sgreben/flagvar"
)

func TestTemplate(t *testing.T) {
	fv := flagvar.Template{}
	var fs flag.FlagSet
	fs.Var(&fv, "template", "")

	err := fs.Parse([]string{"-template", "{{.Abc}}"})
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(fv.Value.Root.String(), template.Must(template.New("").Parse("{{.Abc}}")).Root.String()) {
		t.Fail()
	}
}

func TestTemplateRoot(t *testing.T) {
	fv := flagvar.Template{
		Root: template.New(""),
	}
	var fs flag.FlagSet
	fs.Var(&fv, "template", "")

	err := fs.Parse([]string{"-template", "{{.Abc}}"})
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(fv.Value.Root.String(), template.Must(template.New("").Parse("{{.Abc}}")).Root.String()) {
		t.Fail()
	}
}

func TestTemplateFail(t *testing.T) {
	fv := flagvar.Template{}
	var fs flag.FlagSet
	fs.Var(&fv, "template", "")

	err := fs.Parse([]string{"-template", "{{...}}"})
	if err == nil {
		t.Fail()
	}
}

func TestTemplates(t *testing.T) {
	fv := flagvar.Templates{}
	var fs flag.FlagSet
	fs.Var(&fv, "template", "")

	err := fs.Parse([]string{"-template", "{{.Abc}}"})
	if err != nil {
		t.Fail()
	}
	for _, tmp := range fv.Values {
		if !reflect.DeepEqual(tmp.Root.String(), template.Must(template.New("").Parse("{{.Abc}}")).Root.String()) {
			t.Fail()
		}
	}
}

func TestTemplatesRoot(t *testing.T) {
	fv := flagvar.Templates{
		Root: template.New(""),
	}
	var fs flag.FlagSet
	fs.Var(&fv, "template", "")

	err := fs.Parse([]string{"-template", "{{.Abc}}"})
	if err != nil {
		t.Fail()
	}
	for _, tmp := range fv.Values {
		if !reflect.DeepEqual(tmp.Root.String(), template.Must(template.New("").Parse("{{.Abc}}")).Root.String()) {
			t.Fail()
		}
	}
}

func TestTemplatesFail(t *testing.T) {
	fv := flagvar.Templates{}
	var fs flag.FlagSet
	fs.Var(&fv, "template", "")

	err := fs.Parse([]string{"-template", "{{...}}"})
	if err == nil {
		t.Fail()
	}
}
