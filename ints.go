package flagvar

import (
	"strconv"
	"strings"
)

// Ints is a `flag.Value` for `int` arguments.
// The `Base` and `BitSize` fields are used for parsing when set.
type Ints struct {
	Base    int
	BitSize int

	Values []int64
	Texts  []string
}

// Set is flag.Value.Set
func (fv *Ints) Set(v string) error {
	base := fv.Base
	if base == 0 {
		base = 10
	}
	bitSize := fv.BitSize
	if bitSize == 0 {
		bitSize = 64
	}
	n, err := strconv.ParseInt(v, base, bitSize)
	if err == nil {
		fv.Values = append(fv.Values, n)
		fv.Texts = append(fv.Texts, v)
	}
	return err
}

func (fv *Ints) String() string {
	return strings.Join(fv.Texts, ",")
}
