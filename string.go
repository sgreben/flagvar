package flagvar

import (
	"sort"
	"strings"
)

// Strings is a `flag.Value` for `string` arguments.
type Strings struct {
	Values []string
}

// Set is flag.Value.Set
func (fv *Strings) Set(v string) error {
	fv.Values = append(fv.Values, v)
	return nil
}

func (fv *Strings) String() string {
	return strings.Join(fv.Values, ",")
}

// StringSet is a `flag.Value` for `string` arguments.
// Only distinct values are returned.
type StringSet struct {
	Value map[string]bool
}

// Values returns a string slice of specified values.
func (fv *StringSet) Values() (out []string) {
	for v := range fv.Value {
		out = append(out, v)
	}
	sort.Strings(out)
	return
}

// Set is flag.Value.Set
func (fv *StringSet) Set(v string) error {
	if fv.Value == nil {
		fv.Value = make(map[string]bool)
	}
	fv.Value[v] = true
	return nil
}

func (fv *StringSet) String() string {
	return strings.Join(fv.Values(), ",")
}
