package flagvar

import (
	"encoding/json"
	"strings"
)

type JSON struct {
	Value interface{}
	Text  string
}

// Set is flag.Value.Set
func (fv *JSON) Set(v string) error {
	fv.Text = v
	return json.Unmarshal([]byte(v), &fv.Value)
}

func (fv *JSON) String() string {
	return fv.Text
}

type JSONs struct {
	Values []interface{}
	Texts  []string
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
