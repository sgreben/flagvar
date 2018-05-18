package flagvar

import "os"

type File struct {
	Validate func(os.FileInfo, error) error

	Value string
}

// Set is flag.Value.Set
func (fv *File) Set(v string) error {
	info, err := os.Stat(v)
	if fv.Validate != nil {
		return fv.Validate(info, err)
	}
	return err
}

func (fv *File) String() string {
	return fv.Value
}
