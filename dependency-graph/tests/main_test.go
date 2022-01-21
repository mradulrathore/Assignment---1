package main

import (
	"github.com/mradulrathore/dependency-graph/tests/view"

	"testing"
)

//since we're using bufio for reading input and it has limit on its buffer size, we can not run this integration test.
// To change buffer size:
// info, infoErr := file.Stat()
// var maxSize int
// scanner := bufio.NewScanner(file)
// maxSize = int(info.Size())
// buffer := make([]byte, 0, maxSize)
// scanner.Buffer(buffer, maxSize)
//this test will run properly if we use Scanf for taking input
func TestInit(t *testing.T) {
	view.TestInit(t)
}
