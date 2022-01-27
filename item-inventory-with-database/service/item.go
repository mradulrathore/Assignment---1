package service

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/mradulrathore/item-inventory/service/enum"
)

type Item struct {
	Name     string        `gorm:"column:Name;type:varchar;size:255;" json:"name"`
	Price    float64       `gorm:"column:Price;type:float" json:"price"`
	Quantity int           `gorm:"column:Quantity:type:int;" json:"quantity"`
	Type     enum.ItemType `gorm:"column:Type;type:enum('raw','manufactured','imported');" json:"type"`
}

func (m *Item) TableName() string {
	return "items"
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
	return fmt.Sprintf("\nName: %s\nPrice: %.2f\nQuantity: %d\nType: %s\nTax: %.2f\nCost: %.2f\n", item.Name, item.Price, item.Quantity, item.Type.String(), item.GetTax(), item.GetEffectivePrice())
}

const (
	RAWItmTaxRate                       = 0.125
	ImportDuty                          = 0.100
	FirstImportDuty                     = 100
	SecondImportDuty                    = 200
	FirstImportDutySurchargeAmt         = 5
	SecondImportDutySurchargeAmt        = 10
	ExceedSecondImportDutySurchargeRate = 0.05
	ManufacturedItmTaxRate              = 0.125
	ManufacturedItmExtraTaxRate         = 0.02 //Extra =ItemCost +12.5% Item Cost
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
	if price <= FirstImportDuty {
		return FirstImportDutySurchargeAmt
	} else if price <= SecondImportDuty {
		return SecondImportDutySurchargeAmt
	} else {
		return price * ExceedSecondImportDutySurchargeRate
	}
}
