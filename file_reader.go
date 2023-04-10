package flagvar

import (
	"bufio"
	"os"
)

type FileReader struct {
	Validate func(os.FileInfo, error) error
	Name     string

	Value *bufio.Reader
}

func (fv *FileReader) Set(v string) error {
	_, err := os.Stat(v)
	if err != nil {
		return err
	}

	inFile, err := os.Open(v)
	if err != nil {
		return err
	}

	inputReader := bufio.NewReader(inFile)
	fv.Name = v
	fv.Value = inputReader

	return err
}

func (fv *FileReader) Get() *bufio.Reader {
	return fv.Value
}

func (fv *FileReader) String() string {
	return fv.Name
}
