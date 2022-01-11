package item

// cmd go test -coverprofile=coverage.out

import (
	"testing"
)

func TestAddMoreItems(t *testing.T) {

}

func TestCalculateTaxAndPrice(t *testing.T) {
	var items = []Item{
		// all item details provided
		{
			Name:     "Mango",
			Price:    12,
			Quantity: 1,
			Type:     "raw",
		},
		{
			Name:     "Mango",
			Price:    12,
			Quantity: 1,
			Type:     "manufactured",
		},
		{
			Name:     "Mango",
			Price:    12,
			Quantity: 1,
			Type:     "imported",
		},
		{
			Name:     "Orange",
			Price:    100,
			Quantity: 1,
			Type:     "imported",
		},
		{
			Name:     "Tomato",
			Price:    1000,
			Quantity: 1,
			Type:     "imported",
		},
	}

	expectedSalesTaxLiabilityPerItem := []float64{1.5, 1.77, 6.2, 20, 155}
	for index, item := range items {
		err := item.CalculateTaxAndPrice()
		if err != nil {
			t.Errorf("exception is occuring: %q", err)
		} else if item.SalesTaxLiability != expectedSalesTaxLiabilityPerItem[index] {
			t.Errorf("got %g, wanted %g", item.SalesTaxLiability, expectedSalesTaxLiabilityPerItem[index])
		}

	}
}

func TestGetAllItemDetails(t *testing.T) {
	var items = []Item{
		// all item details provided
		{
			Name:     "Mango",
			Price:    12,
			Quantity: 1,
			Type:     "raw",
		},
		{
			Name:     "Mango",
			Price:    12,
			Quantity: 1,
			Type:     "manufactured",
		},
		{
			Name:     "Mango",
			Price:    12,
			Quantity: 1,
			Type:     "imported",
		},
		{
			Name:     "Orange",
			Price:    100,
			Quantity: 1,
			Type:     "imported",
		},
		{
			Name:     "Tomato",
			Price:    1000,
			Quantity: 1,
			Type:     "imported",
		},
	}

	err := GetAllItemDetails(items)
	if err != nil {
		t.Errorf("exception is occuring: %q", err)
	}
}

func TestValidateItemDetails(t *testing.T) {

	var items = []Item{
		// all item details provided
		{
			Name:     "Mango",
			Price:    100,
			Quantity: 2,
			Type:     "raw",
		},
		// all item details provided
		{
			Name:     "Banana",
			Price:    100,
			Quantity: 2,
			Type:     "manufactured",
		},
		// all item details provided
		{
			Name:     "Orange",
			Price:    100,
			Quantity: 2,
			Type:     "imported",
		},
		// Quantity less than 0
		{
			Name:     "Mango",
			Price:    100,
			Quantity: -2,
			Type:     "raw",
		},
		// type of item not matches predefined type
		{
			Name:     "Mango",
			Price:    100,
			Quantity: -2,
			Type:     "exported",
		},
		// item type missing
		{
			Name:     "Mango",
			Price:    100,
			Quantity: 2,
		},
		// Quantity is not provided and mandatory field(item type) is provided
		{
			Name:  "Mango",
			Price: 100,
			Type:  "raw",
		},
		// Price is not provided and mandatory field(item type) is provided
		{
			Name:     "Mango",
			Quantity: 2,
			Type:     "raw",
		},
		// Name is not provided and mandatory field(item type) is provided
		{
			Price:    100,
			Quantity: 2,
			Type:     "raw",
		},
		// Price less than zero
		{
			Price:    -100,
			Quantity: 2,
			Type:     "raw",
		},
	}

	expected := []bool{true, true, true, false, false, false, true, true, true, false}

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
