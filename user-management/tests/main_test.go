package main

import (
	"os"
	"testing"

	"github.com/mradulrathore/user-management/tests/view"
)

//TODO add .txt test files
//since we're using bufio for reading input and it has limit on its buffer size, we can not run this integration test.
//this test will run properly if we use Scanf for taking input
func TestInit(t *testing.T) {
	view.TestInit(t)
	os.Remove("user-data.json")
}
