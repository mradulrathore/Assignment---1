package item

import (
	"errors"
	"fmt"
	"log"

	"github.com/mradulrathore/onboarding-assignments/constant"
)

type Item struct {
	Name              string
	Price             float64
	Quantity          int
	Type              string
	SalesTaxLiability float64
	FinalPrice        float64
}

func (item *Item) CalculateTaxAndPrice() error {
	switch item.Type {
	case "raw":
		//raw: 12.5% of the item cost
		item.SalesTaxLiability = constant.RAWItmTaxRate * item.Price
		item.FinalPrice = item.Price + item.SalesTaxLiability
	case "manufactured":
		// manufactured: 12.5% of the item cost + 2% of (item cost + 12.5% of the item cost)
		item.SalesTaxLiability = constant.ManufacturedItmTaxRate*item.Price + constant.ManufacturedItmExtraTaxRate*(item.Price+constant.ManufacturedItmTaxRate*item.Price)
		item.FinalPrice = item.Price + item.SalesTaxLiability
	case "imported":
		//imported: 10% import duty on item cost + a surcharge
		item.SalesTaxLiability = constant.ImportDuty * item.Price
		item.FinalPrice = item.Price + item.SalesTaxLiability
		if item.FinalPrice <= constant.ImportDutyLimit1 {
			item.FinalPrice = item.FinalPrice + constant.ImportDutyLimit1SurchargeAmt
			item.SalesTaxLiability = item.SalesTaxLiability + constant.ImportDutyLimit1SurchargeAmt
		} else if item.FinalPrice <= constant.ImportDutyLimit2 {
			item.FinalPrice = item.FinalPrice + constant.ImportDutyLimit2SurchargeAmt
			item.SalesTaxLiability = item.SalesTaxLiability + constant.ImportDutyLimit2SurchargeAmt
		} else {
			item.SalesTaxLiability = item.SalesTaxLiability + item.FinalPrice*constant.ExceedeImportDutyLimit2SurchargeRate
			item.FinalPrice = item.FinalPrice + item.FinalPrice*constant.ExceedeImportDutyLimit2SurchargeRate
		}
	}
	return nil
}

func GetAllItemDetails(items []Item) error {

	for _, item := range items {
		err := item.GetItemDetails()
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func (item Item) GetItemDetails() error {

	fmt.Printf("Item Name: %s \n", item.Name)
	fmt.Printf("Item Price: %g \n", item.Price)
	fmt.Printf("Item Quantity: %d \n", item.Quantity)
	fmt.Printf("Item Type: %s \n", item.Type)
	fmt.Printf("Item Type: %g \n", item.SalesTaxLiability)
	fmt.Printf("Item Type: %g \n \n", item.FinalPrice)

	return nil
}

func (item *Item) SetItemDetails() (bool, error) {

	fmt.Printf("Item Name: ")
	_, err := fmt.Scanf("%s", &(item.Name))
	if err != nil {
		log.Println("scan for Item Name failed, due to ", err)
		return false, err
	}

	fmt.Printf("Item Price: ")
	_, err = fmt.Scanf("%g", &(item.Price))
	if err != nil {
		log.Println("scan for Item Price failed, due to ", err)
		return false, err
	}

	fmt.Printf("Item Quantity: ")
	_, err = fmt.Scanf("%d", &(item.Quantity))
	if err != nil {
		log.Println("scan for Item Quantity failed, due to ", err)
		return false, err
	}

	fmt.Printf("Item Type: ")
	_, err = fmt.Scanf("%s", &(item.Type))
	if err != nil {
		log.Println(" scan for Item type failed, due to ", err)
		return false, err
	}

	ok, err := item.ValidateItemDetails()
	if !ok {
		log.Println(err.Error())
		ok, err = item.SetItemDetails()
		if err != nil {
			log.Println(err)
			return ok, err
		}
	}

	return true, nil

}

func (item *Item) ValidateItemDetails() (bool, error) {
	if len(item.Type) == 0 {
		return false, errors.New("pleae specify item type")
	}
	if item.Quantity < 0 {
		return false, errors.New("quantity can not be negative")
	}
	if item.Price < 0 {
		return false, errors.New("price can not be negative")
	}
	if item.Type != "raw" && item.Type != "manufactured" && item.Type != "imported" {
		return false, errors.New("item type can only be raw, manufactured or imported")
	}
	return true, nil
}
