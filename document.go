package pdf

import (
	"errors"
	"fmt"
	_ "io"
)

// PDF Document Specification Versions
const (
	VERSION_1_0 = iota
	VERSION_1_1
	VERSION_1_2
	VERSION_1_3
	VERSION_1_4
	VERSION_1_5
	VERSION_1_6
	VERSION_1_7
)

const (
	EOL_CR   = "\r"
	EOL_LF   = "\n"
	EOL_CRLF = "\r\n"
)

var (
	ErrorInvalidVersion = errors.New("Invalid version for PDF document")
	ErrorInvalidEol     = errors.New("Invalid end-of-line marker")
)

// PDF Document Model
type Document struct {
	version int
	eol     string
}

// Returns PDF Document specification version
func (d Document) Version() string {
	return fmt.Sprintf("%PDF-1.%d", d.version)
}

func (d Document) Eol() string {
	return d.eol
}

// Sets PDF document specification version
func (d *Document) SetVersion(version int) error {
	var result error
	if version >= 0 && version <= 7 {
		d.version = version
		result = nil
	} else {
		result = ErrorInvalidVersion
	}
	return result
}

// Sets PDF document specification end-of-line marker
func (d *Document) SetEol(eol string) error {
	var result error
	if eol == EOL_CR || eol == EOL_LF || eol == EOL_CRLF {
		d.eol = eol
		result = nil
	} else {
		result = ErrorInvalidEol
	}
	return result
}

// PDF Document Constructor
func New(version int, eol string) (*Document, error) {
	doc := Document{}
	err := doc.SetVersion(version)
	if err != nil {
		return nil, err
	}
	err = doc.SetEol(eol)
	if err != nil {
		return nil, err
	}
	return &doc, nil
}
