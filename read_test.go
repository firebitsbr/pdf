package pdf

import (
	"testing"
)

// Test if error is returned while reading document from file
func TestFromFile_Error(t *testing.T) {
	_, err := FromFile("bad-path%%%^^^")
	if err == nil {
		t.FailNow()
	}
}
