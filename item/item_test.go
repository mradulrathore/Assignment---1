package item

// cmd go test -coverprofile=coverage.out

import (
	"testing"

	"github.com/mradulrathore/onboarding-assignments/constant"
)

func TestAddMoreItems(t *testing.T) {
}

func TestCalculateTaxAndPrice(t *testing.T) {
	var tests = []struct {
		req                       Item
		ExpectedSalesTaxLiability float64
		err                       error
	}{
		// all item details provided
		{
			req: Item{
				Name:     "Mango",
				Price:    12,
				Quantity: 1,
				Type:     constant.Raw,
			},
			ExpectedSalesTaxLiability: 1.5,
			err:                       nil,
		},
		{
			req: Item{
				Name:     "Mango",
				Price:    12,
				Quantity: 1,
				Type:     constant.Manufactured,
			},
			ExpectedSalesTaxLiability: 1.77,
			err:                       nil,
		},
		{
			req: Item{
				Name:     "Mango",
				Price:    12,
				Quantity: 1,
				Type:     constant.Imported,
			},
			ExpectedSalesTaxLiability: 6.2,
			err:                       nil,
		},
		{
			req: Item{
				Name:     "Orange",
				Price:    100,
				Quantity: 1,
				Type:     constant.Imported,
			},
			ExpectedSalesTaxLiability: 20,
			err:                       nil,
		},
		{
			req: Item{
				Name:     "Tomato",
				Price:    1000,
				Quantity: 1,
				Type:     constant.Imported,
			},
			ExpectedSalesTaxLiability: 155,
			err:                       nil,
		},
	}

	for _, tc := range tests {
		err := tc.req.CalculateTaxAndPrice()
		if err != tc.err {
			t.Errorf("got: %v, expected: %v", err, tc.err)
		} else if tc.req.SalesTaxLiability != tc.ExpectedSalesTaxLiability {
			t.Errorf("got %g, wanted %g", tc.req.SalesTaxLiability, tc.ExpectedSalesTaxLiability)
		}

	}

}

func TestGetAllItemDetails(t *testing.T) {

	tests := struct {
		req []Item
		err error
	}{ // all item details provided
		req: []Item{
			{
				Name:     "Mango",
				Price:    100,
				Quantity: 2,
				Type:     constant.Raw,
			},
			{
				Name:     "Orange",
				Price:    100,
				Quantity: 2,
				Type:     constant.Imported,
			},
			{
				Name:     "Orange",
				Price:    100,
				Quantity: 2,
				Type:     "manufactures",
			},
		},
		err: nil,
	}

	err := GetAllItemDetails(tests.req)
	if err != tests.err {
		t.Errorf("exception is occuring: %q", err)
	}
}

func TestValidateItemDetails(t *testing.T) {

	var tests = []struct {
		req Item
		err error
	}{
		// all item details provided
		{
			req: Item{
				Name:     "Mango",
				Price:    100,
				Quantity: 2,
				Type:     constant.Raw,
			},
			err: nil,
		},
		// all item details provided
		{
			req: Item{
				Name:     "Orange",
				Price:    100,
				Quantity: 2,
				Type:     constant.Imported,
			},
			err: nil,
		},
		// Quantity less than 0
		{
			req: Item{
				Name:     "Orange",
				Price:    100,
				Quantity: -2,
				Type:     constant.Imported,
			},
			err: NegativeQuantErr,
		},
		// type of item not matches predefined type
		{
			req: Item{
				Name:     "Mango",
				Price:    100,
				Quantity: 2,
				Type:     "exported",
			},
			err: InvalideItmType,
		},
		// item type missing
		{
			req: Item{
				Name:     "Mango",
				Price:    100,
				Quantity: 2,
			},
			err: InvalideItmType,
		},
		// Quantity is not provided and mandatory field(item type) is provided
		{
			req: Item{
				Name:  "Mango",
				Price: 100,
				Type:  constant.Raw,
			},
			err: nil,
		},
		// Price is not provided and mandatory field(item type) is provided
		{
			req: Item{
				Name:     "Mango",
				Quantity: 2,
				Type:     constant.Raw,
			},
			err: nil,
		},
		// Name is not provided and mandatory field(item type) is provided
		{
			req: Item{
				Price:    100,
				Quantity: 2,
				Type:     constant.Raw,
			},
			err: nil,
		},
		// Price less than zero
		{
			req: Item{
				Price:    -100,
				Quantity: 2,
				Type:     constant.Raw,
			},
			err: NegativePriceErr,
		},
	}

	for _, tc := range tests {
		if err := tc.req.ValidateItemDetails(); err != tc.err {
			t.Errorf("got: %v, expected: %v", err, tc.err)
		}
	}
}
