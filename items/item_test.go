package item

// cmd go test -coverprofile=coverage.out

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	var items = []Item{
		// all item details provided
		{
			Name:     "Mango",
			Price:    12,
			Quantity: 1,
			TypeItem: "raw",
		},
		{
			Name:     "Mango",
			Price:    12,
			Quantity: 1,
			TypeItem: "manufactured",
		},
		{
			Name:     "Mango",
			Price:    12,
			Quantity: 1,
			TypeItem: "imported",
		},
		{
			Name:     "Orange",
			Price:    100,
			Quantity: 1,
			TypeItem: "imported",
		},
		{
			Name:     "Tomato",
			Price:    1000,
			Quantity: 1,
			TypeItem: "imported",
		},
	}

	expectedSalesTaxLiabilityPerItem := []float64{1.5, 1.77, 6.2, 20, 155}
	for index, item := range items {
		err := item.CalculateTax()
		if err != nil {
			t.Errorf("exception is occuring: %q", err)
		} else if item.SalesTaxLiabilityPerItem != expectedSalesTaxLiabilityPerItem[index] {
			t.Errorf("got %g, wanted %g", item.SalesTaxLiabilityPerItem, expectedSalesTaxLiabilityPerItem[index])
		}

	}
}

func TestValidateItemDetails(t *testing.T) {

	var items = []Item{
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
