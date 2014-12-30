package pdf

import (
	"os"
	"path"
	"testing"
)

var (
	TEST_FILEPATH_MINIMAL string
)

func init() {
	TEST_FILEPATH_MINIMAL = path.Join(os.Getenv("GOPATH"), "src", "github.com", "gwaji", "pdf", "minimal.pdf")
}

// Test if error is returned while reading document from file
func TestFromFile_Error(t *testing.T) {
	_, err := FromFile("")
	if err == nil {
		t.FailNow()
	}
}

// Test if there was no error while reading document from file
func TestFromFile_Minimal(t *testing.T) {
	_, err := FromFile(TEST_FILEPATH_MINIMAL)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
}
