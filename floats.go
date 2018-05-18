package flagvar

import (
	"strconv"
	"strings"
)

// Floats is a `flag.Value` for float arguments.
// The `BitSize` field is used for parsing when set.
type Floats struct {
	BitSize int

	Values []float64
	Texts  []string
}

// Set is flag.Value.Set
func (fv *Floats) Set(v string) error {
	bitSize := fv.BitSize
	if bitSize == 0 {
		bitSize = 64
	}
	n, err := strconv.ParseFloat(v, bitSize)
	if err == nil {
		fv.Values = append(fv.Values, n)
		fv.Texts = append(fv.Texts, v)
	}
	return err
}

func (fv *Floats) String() string {
	return strings.Join(fv.Texts, ",")
}
