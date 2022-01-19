package main

import (
	"os"
	"testing"

	"github.com/mradulrathore/user-management/tests/view"
)

//TODO create test files
func TestInit(t *testing.T) {

	view.TestInit(t)
	os.Remove("user-data.json")
}
