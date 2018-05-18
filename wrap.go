package flagvar

import "flag"

type WrapFunc func() flag.Value

// Set is flag.Value.Set
func (fv WrapFunc) Set(v string) error {
	return fv().Set(v)
}

func (fv WrapFunc) String() string {
	return fv().String()
}

type Wrap struct {
	Value   flag.Value
	Updated func()
}

// Set is flag.Value.Set
func (fv *Wrap) Set(v string) error {
	err := fv.Value.Set(v)
	if err == nil {
		fv.Updated()
	}
	return err
}

func (fv *Wrap) String() string {
	return fv.Value.String()
}
