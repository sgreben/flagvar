package flagvar

import (
	"regexp"
	"strings"
)

// Regexp is a `flag.Value` for regular expression arguments.
type Regexp struct {
	Value *regexp.Regexp
	Text  string
}

// Set is flag.Value.Set
func (fv *Regexp) Set(v string) error {
	re, err := regexp.Compile(v)
	if err != nil {
		return err
	}
	fv.Text = v
	fv.Value = re
	return nil
}

func (fv *Regexp) String() string {
	return fv.Text
}

// Regexps is a `flag.Value` for regular expression arguments.
type Regexps struct {
	Values []*regexp.Regexp
	Texts  []string
}

// Set is flag.Value.Set
func (fv *Regexps) Set(v string) error {
	re, err := regexp.Compile(v)
	if err != nil {
		return err
	}
	fv.Texts = append(fv.Texts, v)
	fv.Values = append(fv.Values, re)
	return nil
}

func (fv *Regexps) String() string {
	return strings.Join(fv.Texts, ",")
}
