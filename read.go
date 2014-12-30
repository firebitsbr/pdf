package pdf

import (
	"bytes"
	"io/ioutil"
)

// Read PDF Document from file and return
// Document instance and error if needed
func FromFile(path string) (*Document, error) {
	stream, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewReader(stream)

	// Scanning document (byte stream)
	var (
		current_position uint // Current position in byte stream
		b                byte
	)
	for err == nil && (b != EOL_CR[0] || b != EOL_LF[0]) {
		b, err = buffer.ReadByte()

		current_position++
	}

	return New(VERSION_1_0, EOL_LF)
}
