package item

// cmd go test -coverprofile=coverage.out

import (
	"errors"
	"testing"

	"github.com/mradulrathore/onboarding-assignments/item-inventory/item/enum"
)

//Doubt
func TestNew(t *testing.T) {
	_, invalidTypeErr := enum.ItemTypeString("exported")

	var tests = []struct {
		scenario string
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
			err:      nil,
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
			err:      errors.New("negative value"),
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
			err:      errors.New("negative value"),
		},
	}

	for _, tc := range tests {
		_, err := New(tc.name, tc.price, tc.quantity, tc.typeItem)
		if err != tc.err {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}
