package flagvar

import (
	"encoding/json"
	"strings"
)

// JSON is a `flag.Value` for JSON arguments.
type JSON struct {
	Value interface{}
	Text  string
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *JSON) Help() string {
	return "a JSON value"
}

// Set is flag.Value.Set
func (fv *JSON) Set(v string) error {
	fv.Text = v
	return json.Unmarshal([]byte(v), &fv.Value)
}

func (fv *JSON) String() string {
	return fv.Text
}

// JSONs is a `flag.Value` for JSON arguments.
type JSONs struct {
	Values []interface{}
	Texts  []string
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *JSONs) Help() string {
	return "a JSON value"
}

// Set is flag.Value.Set
func (fv *JSONs) Set(v string) error {
	var value interface{}
	err := json.Unmarshal([]byte(v), &value)
	if err == nil {
		fv.Texts = append(fv.Texts, v)
		fv.Values = append(fv.Values, value)
	}
	return err
}

func (fv *JSONs) String() string {
	return strings.Join(fv.Texts, ",")
}
