package item

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/mradulrathore/onboarding-assignments/item-inventory/item/enum"
)

type Item struct {
	Name     string        `json:"name"`
	Price    float64       `json:"price"`
	Quantity int           `json:"quantity"`
	Type     enum.ItemType `json:"type"`
}

func New(name string, price float64, quantity int, typeItem string) (Item, error) {
	var item Item
	var err error
	item.Name = name
	item.Price = price
	item.Quantity = quantity
	item.Type, err = enum.ItemTypeString(typeItem)
	if err != nil {
		return Item{}, err
	}

	if err = item.validate(); err != nil {
		return Item{}, err
	}
	return item, nil
}

func (item Item) validate() error {
	return validation.ValidateStruct(&item,
		validation.Field(&item.Quantity, validation.By(checkNegativeValue)),
		validation.Field(&item.Price, validation.By(checkNegativeValue)),
	)
}

func checkNegativeValue(value interface{}) error {
	err := fmt.Errorf("%v", "negative value")
	switch t := value.(type) {
	case int:
		if t < 0 {
			return err
		}
	case float64:
		if t < 0.0 {
			return err
		}
	}

	return nil
}

func (item Item) Invoice() string {
	return fmt.Sprintf("[%s, %f, %d,%s,%.2f,%.2f]", item.Name, item.Price, item.Quantity, item.Type.String(), item.GetTax(), item.GetEffectivePrice())
}

const (
	RAWItmTaxRate                        = 0.125
	ImportDuty                           = 0.100
	ImportDutyLimit1                     = 100
	ImportDutyLimit2                     = 200
	ImportDutyLimit1SurchargeAmt         = 5
	ImportDutyLimit2SurchargeAmt         = 10
	ExceedeImportDutyLimit2SurchargeRate = 0.05
	ManufacturedItmTaxRate               = 0.125
	ManufacturedItmExtraTaxRate          = 0.02 //Extra =ItemCost +12.5% Item Cost
)

func (item Item) GetTax() float64 {
	var tax float64
	switch item.Type {
	case enum.Raw:
		//raw: 12.5% of the item cost
		tax = RAWItmTaxRate * item.Price
	case enum.Manufactured:
		// manufactured: 12.5% of the item cost + 2% of (item cost + 12.5% of the item cost)
		tax = ManufacturedItmTaxRate*item.Price + ManufacturedItmExtraTaxRate*(item.Price+ManufacturedItmTaxRate*item.Price)
	case enum.Imported:
		//imported: 10% import duty on item cost
		tax = ImportDuty * item.Price
	}

	return tax
}

func (item Item) GetEffectivePrice() float64 {
	var effectivePrice float64
	surcharge := 0.0
	tax := item.GetTax()

	switch item.Type {
	case enum.Raw:
		effectivePrice = item.Price + tax + surcharge
	case enum.Manufactured:
		effectivePrice = item.Price + tax + surcharge
	case enum.Imported:
		priceTemp := item.Price + tax
		surcharge = item.importSurcharge(priceTemp)
		effectivePrice = priceTemp + surcharge
	}

	return effectivePrice
}

func (item Item) importSurcharge(price float64) float64 {
	if price <= ImportDutyLimit1 {
		return ImportDutyLimit1SurchargeAmt
	} else if price <= ImportDutyLimit2 {
		return ImportDutyLimit2SurchargeAmt
	} else {
		return price * ExceedeImportDutyLimit2SurchargeRate
	}
}
