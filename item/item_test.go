package item

// cmd go test -coverprofile=coverage.out

import (
	"testing"

	"github.com/mradulrathore/onboarding-assignments/item/enum"
)

// // func TestNew(t *testing.T) {

// // 	var tests = []struct {
// // 		req Item
// // 		tax float32
// // 		err error
// // 	}{
// // 		// all item details provided
// // 		{
// // 			req: Item{
// // 				Name:     "Mango",
// // 				Price:    12,
// // 				Quantity: 1,
// // 				Type:     Raw,
// // 			},
// // 			tax: 1.5,
// // 			err: nil,
// // 		},
// // 		{
// // 			req: Item{
// // 				Name:     "Mango",
// // 				Price:    12,
// // 				Quantity: 1,
// // 				Type:     Manufactured,
// // 			},
// // 			tax: 1.77,
// // 			err: nil,
// // 		},
// // 		{
// // 			req: Item{
// // 				Name:     "Mango",
// // 				Price:    12,
// // 				Quantity: 1,
// // 				Type:     Imported,
// // 			},
// // 			tax: 6.2,
// // 			err: nil,
// // 		},
// // 		{
// // 			req: Item{
// // 				Name:     "Orange",
// // 				Price:    100,
// // 				Quantity: 1,
// // 				Type:     Imported,
// // 			},
// // 			tax: 20,
// // 			err: nil,
// // 		},
// // 		{
// // 			req: Item{
// // 				Name:     "Tomato",
// // 				Price:    1000,
// // 				Quantity: 1,
// // 				Type:     Imported,
// // 			},
// // 			tax: 155,
// // 			err: nil,
// // 		},
// // 	}

// // 	for _, tc := range tests {
// // 		item, err := New(tc.req.Name, tc.req.Price, tc.req.Quantity, tc.req.Type)
// // 		if err != tc.err {
// // 			t.Errorf("got: %v, expected: %v", err, tc.err)
// // 		} else if tc.tax != tc.ExpectedSalesTaxLiability {
// // 			t.Errorf("got %g, wanted %g", tc.req.SalesTaxLiability, tc.ExpectedSalesTaxLiability)
// // 		}

// // 	}

// // }

func TestNew(t *testing.T) {
	_, emptyTypeErr := enum.ItemTypeString("")
	_, invalidTypeErr := enum.ItemTypeString("exported")

	var tests = []struct {
		scenario string
		desc     string
		name     string
		price    float64
		quantity int
		typeItem string
		err      error
	}{
		{
			scenario: "all item details provided",
			name:     "Mango",
			price:    100,
			quantity: 2,
			typeItem: "raw",
			err:      nil,
		},
		{
			scenario: "all item details provided",
			name:     "Orange",
			price:    100,
			quantity: 2,
			typeItem: "imported",
			err:      nil,
		},
		{
			scenario: "quantity less than 0",
			name:     "Orange",
			price:    100,
			quantity: -2,
			typeItem: "imported",
			err:      NegativeQuantErr,
		},
		{
			scenario: "type of item not matches predefined type",
			name:     "Mango",
			price:    100,
			quantity: 2,
			typeItem: "exported",
			err:      invalidTypeErr,
		},
		{
			scenario: "item type not provided",
			name:     "Mango",
			price:    100,
			quantity: 2,
			err:      emptyTypeErr,
		},
		{
			scenario: "quantity is not provided and mandatory field(item type) is provided",
			name:     "Mango",
			price:    100,
			typeItem: "raw",
			err:      nil,
		},
		{
			scenario: "price is not provided and mandatory field(item type) is provided",
			name:     "Mango",
			quantity: 2,
			typeItem: "raw",
			err:      nil,
		},
		{
			scenario: "name is not provided and mandatory field(item type) is provided",
			price:    100,
			quantity: 2,
			typeItem: "raw",
			err:      nil,
		},
		{
			scenario: "price less than zero",
			price:    -100,
			quantity: 2,
			typeItem: "raw",
			err:      NegativePriceErr,
		},
	}

	for _, tc := range tests {
		_, err := New(tc.name, tc.price, tc.quantity, tc.typeItem)
		if err != tc.err {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}
