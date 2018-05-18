package flagvar

import (
	"net/url"
	"strings"
)

// URL is a `flag.Value` for `url.URL` arguments.
type URL struct {
	Value *url.URL
	Text  string
}

// Set is flag.Value.Set
func (fv *URL) Set(v string) error {
	u, err := url.Parse(v)
	if err == nil {
		fv.Text = v
		fv.Value = u
	}
	return err
}

func (fv *URL) String() string {
	return fv.Text
}

// URLs is a `flag.Value` for `url.URL` arguments.
type URLs struct {
	Values []*url.URL
	Texts  []string
}

// Set is flag.Value.Set
func (fv *URLs) Set(v string) error {
	u, err := url.Parse(v)
	if err == nil {
		fv.Texts = append(fv.Texts, v)
		fv.Values = append(fv.Values, u)
	}
	return err
}

func (fv *URLs) String() string {
	return strings.Join(fv.Texts, ",")
}
