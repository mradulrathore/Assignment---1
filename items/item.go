package item

import (
	constant "application/constants"
	"errors"
	"fmt"
	"log"
)

type Item struct {
	Name                     string
	Price                    float64
	Quantity                 int
	TypeItem                 string
	SalesTaxLiabilityPerItem float64
	FinalPrice               float64
}

func (item *Item) CalculateTaxAndPrice() error {
	switch item.TypeItem {
	case "raw":
		//raw: 12.5% of the item cost
		item.SalesTaxLiabilityPerItem = constant.RAWItmTaxRate * item.Price / 100
		item.FinalPrice = item.Price + item.SalesTaxLiabilityPerItem
	case "manufactured":
		// manufactured: 12.5% of the item cost + 2% of (item cost + 12.5% of the item cost)
		item.SalesTaxLiabilityPerItem = constant.ManufacturedItmTaxRate*item.Price/100 + constant.ManufacturedItmExtraTaxRate*(item.Price+constant.ManufacturedItmTaxRate*item.Price/100)/100
		item.FinalPrice = item.Price + item.SalesTaxLiabilityPerItem
	case "imported":
		//imported: 10% import duty on item cost + a surcharge
		item.SalesTaxLiabilityPerItem = constant.ImportDuty * item.Price / 100
		item.FinalPrice = item.Price + item.SalesTaxLiabilityPerItem
		if item.FinalPrice <= constant.ImportDutyLimit1 {
			item.FinalPrice = item.FinalPrice + constant.ImportDutyLimit1SurchargeAmt
			item.SalesTaxLiabilityPerItem = item.SalesTaxLiabilityPerItem + constant.ImportDutyLimit1SurchargeAmt
		} else if item.FinalPrice <= constant.ImportDutyLimit2 {
			item.FinalPrice = item.FinalPrice + constant.ImportDutyLimit2SurchargeAmt
			item.SalesTaxLiabilityPerItem = item.SalesTaxLiabilityPerItem + constant.ImportDutyLimit2SurchargeAmt
		} else {
			item.SalesTaxLiabilityPerItem = item.SalesTaxLiabilityPerItem + item.FinalPrice*constant.ExceedeImportDutyLimit2SurchargeAmt/100
			item.FinalPrice = item.FinalPrice + item.FinalPrice*constant.ExceedeImportDutyLimit2SurchargeAmt/100
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
	fmt.Printf("Item Type: %s \n", item.TypeItem)
	fmt.Printf("Item Type: %g \n", item.SalesTaxLiabilityPerItem)
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
	_, err = fmt.Scanf("%s", &(item.TypeItem))
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
	if len(item.TypeItem) == 0 {
		return false, errors.New("pleae specify item type")
	}
	if item.Quantity < 0 {
		return false, errors.New("quantity can not be negative")
	}
	if item.Price < 0 {
		return false, errors.New("price can not be negative")
	}
	if item.TypeItem != "raw" && item.TypeItem != "manufactured" && item.TypeItem != "imported" {
		return false, errors.New("item type can only be raw, manufactured or imported")
	}
	return true, nil
}

func GetUserChoice() (string, error) {

	fmt.Println("Do you want to enter details of any other item (" + constant.Accept + "/" + constant.Deny + ")")
	var moreItems string = constant.Accept
	_, err := fmt.Scanf("%s", &moreItems)
	if err != nil {
		log.Println(err)
		return "", err
	}

	_, err = ValidateConfirmation(moreItems)

	for err != nil {

		_, err = fmt.Scanf("%s", &moreItems)
		if err != nil {
			log.Println(err)
			return "", err
		}
		_, err = ValidateConfirmation(moreItems)
	}

	return moreItems, nil
}

// validate whether userChoice is eiter Accept or Deny
func ValidateConfirmation(userChoice string) (bool, error) {

	if userChoice != constant.Accept && userChoice != constant.Deny {
		log.Println("enter either " + constant.Accept + " or " + constant.Deny)
		return false, errors.New("enter either " + constant.Accept + " or " + constant.Deny)
	}

	return true, nil
}
