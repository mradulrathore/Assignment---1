package main

import (
	"os"
	"testing"

	"github.com/mradulrathore/user-management/tests/view"
)

func TestInit(t *testing.T) {

	view.TestInit(t)
	os.Remove("user-data.json")
}
