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

// func checkError(msg string, err error) {
// 	if err != nil {
// 		log.Print(msg, err)
// 	}
// }

func (item *Item) CalculateTax() error {
	switch item.TypeItem {
	case "raw":
		item.SalesTaxLiabilityPerItem = constant.TaxRateForRAW * item.Price / 100
		item.FinalPrice = item.Price + item.SalesTaxLiabilityPerItem
	case "manufactured":
		item.SalesTaxLiabilityPerItem = constant.TaxRateForManufacturedItemOnItemCost*item.Price/100 + constant.TaxRateForManufactureItemOnCombined*(item.Price+constant.TaxRateForManufacturedItemOnItemCost*item.Price/100)/100
		item.FinalPrice = item.Price + item.SalesTaxLiabilityPerItem
	case "imported":
		item.SalesTaxLiabilityPerItem = constant.ImportDuty * item.Price / 100
		item.FinalPrice = item.Price + item.SalesTaxLiabilityPerItem
		if item.FinalPrice <= constant.ImportDutyLimit1 {
			item.FinalPrice = item.FinalPrice + constant.SurchargeAmountForFinalCostUptoImportDutyLimit1
			item.SalesTaxLiabilityPerItem = item.SalesTaxLiabilityPerItem + constant.SurchargeAmountForFinalCostUptoImportDutyLimit1
		} else if item.FinalPrice <= constant.ImportDutyLimit2 {
			item.FinalPrice = item.FinalPrice + constant.SurchargeAmountForFinalCostUptoImportDutyLimit2
			item.SalesTaxLiabilityPerItem = item.SalesTaxLiabilityPerItem + constant.SurchargeAmountForFinalCostUptoImportDutyLimit2
		} else {
			item.SalesTaxLiabilityPerItem = item.SalesTaxLiabilityPerItem + item.FinalPrice*constant.SurchargeRateForFinalCostExceedeImportDutyLimit2/100
			item.FinalPrice = item.FinalPrice + item.FinalPrice*constant.SurchargeRateForFinalCostExceedeImportDutyLimit2/100

		}

	}
	return nil
}

func GetAllItemDetails(items []Item) {

	for _, item := range items {
		item.GetItemDetails()
	}
}

func (item Item) GetItemDetails() {

	fmt.Printf("Item Name: %s \n", item.Name)
	fmt.Printf("Item Price: %g \n", item.Price)
	fmt.Printf("Item Quantity: %d \n", item.Quantity)
	fmt.Printf("Item Type: %s \n", item.TypeItem)
	fmt.Printf("Item Type: %g \n", item.SalesTaxLiabilityPerItem)
	fmt.Printf("Item Type: %g \n \n", item.FinalPrice)
}

func (item *Item) SetItemDetails() (bool, error) {

	fmt.Printf("Item Name: ")
	// MANGO JUICE
	_, err := fmt.Scanf("%s", &(item.Name))
	if err != nil {
		log.Print("  Scan for Item Name failed, due to ", err)
		return false, err
	}

	fmt.Printf("Item Price: ")
	_, err = fmt.Scanf("%g", &(item.Price))
	if err != nil {
		log.Print("  Scan for Item Price failed, due to ", err)
		return false, err
	}

	fmt.Printf("Item Quantity: ")
	_, err = fmt.Scanf("%d", &(item.Quantity))
	if err != nil {
		log.Print("  Scan for Item Quantity failed, due to ", err)
		return false, err
	}

	fmt.Printf("Item Type: ")
	_, err = fmt.Scanf("%s", &(item.TypeItem))
	if err != nil {
		log.Print("  Scan for Item type failed, due to ", err)
		return false, err
	}

	ok, err := item.ValidateItemDetails()
	if !ok {
		fmt.Println(err.Error())
		ok, err = item.SetItemDetails()
		if err != nil {
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

func AddMoreItems() (string, error) {
	fmt.Println("Do you want to enter details of any other item (" + constant.Accept + "/" + constant.Deny + ")")
	var moreItems string = constant.Accept
	_, err := fmt.Scanf("%s", &moreItems)
	if err != nil {
		return "", err
	}

	err = ValidateConfirmation(moreItems)

	for err != nil {
		fmt.Println(err.Error())
		_, err = fmt.Scanf("%s", &moreItems)
		if err != nil {
			return "", err
		}
		err = ValidateConfirmation(moreItems)
	}

	return moreItems, nil

}

func ValidateConfirmation(userChoice string) error {

	if userChoice != constant.Accept && userChoice != constant.Deny {
		return errors.New("enter either " + constant.Accept + " or " + constant.Deny)
	}

	return nil
}
