package flagvar

import "flag"

// Alternative tries to parse the argument using `Either`, and if that fails, using `Or`.
// `EitherOk` is true if the first attempt succeed.
type Alternative struct {
	Either   flag.Value
	Or       flag.Value
	EitherOk bool
}

// Set is flag.Value.Set
func (fv *Alternative) Set(v string) error {
	err := fv.Either.Set(v)
	fv.EitherOk = err == nil
	if err != nil {
		return fv.Or.Set(v)
	}
	return nil
}

func (fv *Alternative) String() string {
	if fv.EitherOk {
		return fv.Either.String()
	}
	if fv.Or != nil {
		return fv.Or.String()
	}
	return ""
}
