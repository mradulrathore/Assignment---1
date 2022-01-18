package view

// cmd go test -coverprofile=coverage.out

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

type testStr struct {
	name       string
	price      float64
	quantity   int
	typeItem   string
	userChoice string
}

func TestInitialize(t *testing.T) {
	testRaw := testStr{
		name:       "Mango",
		price:      100,
		quantity:   2,
		typeItem:   "raw",
		userChoice: "n",
	}
	testManufactured := testStr{
		name:       "Mango",
		price:      100,
		quantity:   2,
		typeItem:   "manufactured",
		userChoice: "n",
	}
	testImported := testStr{
		name:       "Mango",
		price:      100,
		quantity:   2,
		typeItem:   "imported",
		userChoice: "n",
	}

	testInvalidItmType := testStr{
		name:       "Mango",
		price:      100,
		quantity:   2,
		typeItem:   "exported",
		userChoice: "n",
	}

	tests := []struct {
		scenario string
		req      *os.File
		err      error
	}{{
		scenario: "all item details provided for raw",
		req:      setInput(testRaw),
		err:      nil,
	}, {
		scenario: "all item details provided for manufactured",
		req:      setInput(testManufactured),
		err:      nil,
	}, {
		scenario: "all item details provided for imported",
		req:      setInput(testImported),
		err:      nil,
	}, {
		scenario: "invalid item type, exported",
		req:      setInput(testInvalidItmType),
		err:      errors.New("invalide type error"),
	},
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	for _, tc := range tests {
		os.Stdin = tc.req
		err := Initialize()
		if err != nil && tc.err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if err == nil && tc.err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}

		if err := tc.req.Close(); err != nil {
			log.Fatal(err)
		}
	}

}

func setInput(test testStr) *os.File {
	content := fmt.Sprintf("%s\n%g %d\n%s\n%s", test.name, test.price, test.quantity, test.typeItem, test.userChoice)
	contentB := []byte(content)

	tmpfile, err := ioutil.TempFile("", "temp")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(contentB); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	return tmpfile
}
