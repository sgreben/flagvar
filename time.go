package flagvar

import (
	"strings"
	"time"
)

type Time struct {
	Layout string

	Value time.Time
	Text  string
}

// Set is flag.Value.Set
func (fv *Time) Set(v string) error {
	t, err := time.Parse(fv.Layout, v)
	if err == nil {
		fv.Text = v
		fv.Value = t
	}
	return err
}

func (fv *Time) String() string {
	return fv.Text
}

type Times struct {
	Layout string

	Values []time.Time
	Texts  []string
}

// Set is flag.Value.Set
func (fv *Times) Set(v string) error {
	t, err := time.Parse(fv.Layout, v)
	if err == nil {
		fv.Texts = append(fv.Texts, v)
		fv.Values = append(fv.Values, t)
	}
	return err
}

func (fv *Times) String() string {
	return strings.Join(fv.Texts, ",")
}
