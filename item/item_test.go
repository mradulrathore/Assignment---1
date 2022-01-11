package item

// // cmd go test -coverprofile=coverage.out

// import (
// 	"testing"

// 	"github.com/mradulrathore/onboarding-assignments/item/enum"
// )

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

// func TestNew(t *testing.T) {

// 	var tests = []struct {
// 		desc string
// 		req  Item
// 		err  error
// 	}{
// 		// all item details provided
// 		{
// 			req: Item{
// 				Name:     "Mango",
// 				Price:    100,
// 				Quantity: 2,
// 				Type:     enum.ItemTypeString("raw"),
// 			},
// 			err: nil,
// 		},
// 		// all item details provided
// 		{
// 			req: Item{
// 				Name:     "Orange",
// 				Price:    100,
// 				Quantity: 2,
// 				Type:     "imported",
// 			},
// 			err: nil,
// 		},
// 		// Quantity less than 0
// 		{
// 			req: Item{
// 				Name:     "Orange",
// 				Price:    100,
// 				Quantity: -2,
// 				Type:     Imported,
// 			},
// 			err: NegativeQuantErr,
// 		},
// 		// type of item not matches predefined type
// 		{
// 			req: Item{
// 				Name:     "Mango",
// 				Price:    100,
// 				Quantity: 2,
// 				Type:     "exported",
// 			},
// 			err: InvalideItmType,
// 		},
// 		// item type missing
// 		{
// 			req: Item{
// 				Name:     "Mango",
// 				Price:    100,
// 				Quantity: 2,
// 			},
// 			err: InvalideItmType,
// 		},
// 		// Quantity is not provided and mandatory field(item type) is provided
// 		{
// 			req: Item{
// 				Name:  "Mango",
// 				Price: 100,
// 				Type:  Raw,
// 			},
// 			err: nil,
// 		},
// 		// Price is not provided and mandatory field(item type) is provided
// 		{
// 			req: Item{
// 				Name:     "Mango",
// 				Quantity: 2,
// 				Type:     Raw,
// 			},
// 			err: nil,
// 		},
// 		// Name is not provided and mandatory field(item type) is provided
// 		{
// 			req: Item{
// 				Price:    100,
// 				Quantity: 2,
// 				Type:     Raw,
// 			},
// 			err: nil,
// 		},
// 		// Price less than zero
// 		{
// 			req: Item{
// 				Price:    -100,
// 				Quantity: 2,
// 				Type:     Raw,
// 			},
// 			err: NegativePriceErr,
// 		},
// 	}

// 	for _, tc := range tests {
// 		_, err := New(tc.req.Name, tc.req.Price, tc.req.Quantity, tc.req.Type)
// 		if err != tc.err {
// 			t.Errorf("got: %v, expected: %v", err, tc.err)
// 		}
// 	}
// }
