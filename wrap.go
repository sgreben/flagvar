package flagvar

import "flag"

// WrapPointer wraps a pointer to a `flag.Value`
// This can be used to switch between different argument parsers.
type WrapPointer struct {
	Value *flag.Value
}

// Set is flag.Value.Set
func (fv *WrapPointer) Set(v string) error {
	return (*fv.Value).Set(v)
}

func (fv WrapPointer) String() string {
	return (*fv.Value).String()
}

// WrapFunc wraps a nullary function returning a `flag.Value`
// This can be used to switch between different argument parsers.
type WrapFunc func() flag.Value

// Set is flag.Value.Set
func (fv WrapFunc) Set(v string) error {
	return fv().Set(v)
}

func (fv WrapFunc) String() string {
	return fv().String()
}

// Wrap wraps a `flag.Value` and calls `Updated` each time the underlying value is set.
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
