package flagvar

import (
	"fmt"
	"strings"
)

type Enum struct {
	Choices []string

	Value string
}

// Set is flag.Value.Set
func (fv *Enum) Set(v string) error {
	for _, c := range fv.Choices {
		if strings.EqualFold(c, v) {
			fv.Value = c
			return nil
		}
	}
	return fmt.Errorf(`"%s" must be one of [%s]`, v, strings.Join(fv.Choices, " "))
}

func (fv *Enum) String() string {
	return fv.Value
}

type Enums struct {
	Choices []string

	Values []string
}

// Set is flag.Value.Set
func (fv *Enums) Set(v string) error {
	for _, c := range fv.Choices {
		if strings.EqualFold(c, v) {
			fv.Values = append(fv.Values, c)
			return nil
		}
	}
	return fmt.Errorf(`"%s" must be one of [%s]`, v, strings.Join(fv.Choices, " "))
}

func (fv *Enums) String() string {
	return strings.Join(fv.Values, ",")
}

type EnumSet struct {
	Choices []string

	Value map[string]bool
}

// Values returns a string slice of specified values.
func (fv *EnumSet) Values() (out []string) {
	for v := range fv.Value {
		out = append(out, v)
	}
	return
}

// Set is flag.Value.Set
func (fv *EnumSet) Set(v string) error {
	var ok bool
	for _, c := range fv.Choices {
		if strings.EqualFold(c, v) {
			v = c
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf(`"%s" must be one of [%s]`, v, strings.Join(fv.Choices, " "))
	}
	if fv.Value == nil {
		fv.Value = make(map[string]bool)
	}
	fv.Value[v] = true
	return nil
}

func (fv *EnumSet) String() string {
	return strings.Join(fv.Values(), ",")
}
