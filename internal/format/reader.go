package format

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

type Reader struct {
	io.Reader
	Format
}

func NewReader(reader io.Reader, format Format) Reader {
	return Reader{
		Reader: reader,
		Format: format,
	}
}

func (r Reader) Read(v interface{}) error {
	switch r.Format {
	case JSON:
		return r.readJson(v)
	default:
		return fmt.Errorf("format not supported")
	}
}

func (r Reader) readJson(v interface{}) error {
	b, err := ioutil.ReadAll(r.Reader)
	if err != nil {
		// If there was an error reading the file, exit
		return err
	}
	err = json.Unmarshal(b, v)
	if err != nil {
		// If there was an error parsing it, exit
		return err
	}
	return nil
}
