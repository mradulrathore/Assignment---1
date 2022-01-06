package main

// cmd go test -coverprofile=coverage.out

import (
	itemFile "application/items"
	"testing"
)

func TestValidateItemDetails(t *testing.T) {

	var items = []itemFile.Item{
		// all item details provided
		{
			Name:     "Mango",
			Price:    100,
			Quantity: 2,
			TypeItem: "raw",
		},
		// all item details provided
		{
			Name:     "Banana",
			Price:    100,
			Quantity: 2,
			TypeItem: "manufactured",
		},
		// all item details provided
		{
			Name:     "Orange",
			Price:    100,
			Quantity: 2,
			TypeItem: "imported",
		},
		// Quantity less than 0
		{
			Name:     "Mango",
			Price:    100,
			Quantity: -2,
			TypeItem: "raw",
		},
		// type of item not matches predefined type
		{
			Name:     "Mango",
			Price:    100,
			Quantity: -2,
			TypeItem: "exported",
		},
		// item type missing
		{
			Name:     "Mango",
			Price:    100,
			Quantity: 2,
		},
		// Quantity is not provided and mandatory field(item type) is provided
		{
			Name:     "Mango",
			Price:    100,
			TypeItem: "raw",
		},
		// Price is not provided and mandatory field(item type) is provided
		{
			Name:     "Mango",
			Quantity: 2,
			TypeItem: "raw",
		},
		// Name is not provided and mandatory field(item type) is provided
		{
			Price:    100,
			Quantity: 2,
			TypeItem: "raw",
		},
	}

	expected := []bool{true, true, true, false, false, false, true, true, true}

	for index, item := range items {
		if ok, err := item.ValidateItemDetails(); ok != expected[index] {
			if !expected[index] {
				t.Errorf("exception is occuring: %q", err)
			} else {
				t.Errorf("exception is not occuring: %q", err)
			}

		}
	}
}
