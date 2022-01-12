package view

// cmd go test -coverprofile=coverage.out

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestInitialize(t *testing.T) {

	tests := struct {
		scenario string
		req      *os.File
		err      error
	}{
		scenario: "all item details provided",
		req:      setInput(),
		err:      nil,
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tests.req

	if err := Initialize(); err != tests.err {
		t.Errorf("userInput failed: %v", err)
	}

	if err := tests.req.Close(); err != nil {
		log.Fatal(err)
	}
}

func setInput() *os.File {

	test := struct {
		Name       string
		Price      float64
		Quantity   int
		TypeItem   string
		userChoice string
	}{
		Name:       "Mango\n",
		Price:      100,
		Quantity:   2,
		TypeItem:   "raw\n",
		userChoice: "n\n",
	}

	content := fmt.Sprintf("%s\n%g %d\n%s\n%s", test.Name, test.Price, test.Quantity, test.TypeItem, test.userChoice)
	contentB := []byte(content)

	tmpfile, err := ioutil.TempFile("", "temp")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(contentB); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	return tmpfile
}

// func TestValidateConfirmation(t *testing.T) {

// 	var tests = []struct {
// 		userChoice string
// 		err        error
// 	}{
// 		{userChoice: "y", err: nil},
// 		{userChoice: "n", err: nil},
// 		{userChoice: "er", err: InvalidUsrChoice},
// 		{userChoice: "yes", err: InvalidUsrChoice},
// 		{userChoice: "no", err: InvalidUsrChoice},
// 	}

// 	for _, tc := range tests {
// 		if err := ValidateConfirmation(tc.userChoice); err != tc.err {
// 			t.Errorf("got: %v, expected: %v", err, tc.err)
// 		}
// 	}
// }
