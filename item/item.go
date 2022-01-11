package item

import (
	"fmt"
	"log"

	custErr "github.com/mradulrathore/onboarding-assignments/error"

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

func (item *Item) CalculateRAWTax() error {
	item.SalesTaxLiability = constant.RAWItmTaxRate * item.Price
	return nil
}

func (item *Item) CalculateRAWFinalPrice() error {
	item.FinalPrice = item.Price + item.SalesTaxLiability
	return nil
}

func (item *Item) CalculateManufcaturedTax() error {
	item.SalesTaxLiability = constant.ManufacturedItmTaxRate*item.Price + constant.ManufacturedItmExtraTaxRate*(item.Price+constant.ManufacturedItmTaxRate*item.Price)
	return nil
}

func (item *Item) CalculateManufacturedFinalPrice() error {
	item.FinalPrice = item.Price + item.SalesTaxLiability
	return nil
}

func (item *Item) CalculateImportTax() error {
	item.SalesTaxLiability = constant.ImportDuty * item.Price
	return nil
}

func (item *Item) CalculateImportFinalPrice() error {
	item.FinalPrice = item.Price + item.SalesTaxLiability
	return nil
}

func (item *Item) ApplyImportSurcharge() error {
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

	return nil
}

func (item *Item) CalculateTaxAndPrice() error {
	switch item.Type {
	case constant.Raw:
		//raw: 12.5% of the item cost
		err := item.CalculateRAWTax()
		if err != nil {
			return err
		}

		err = item.CalculateRAWFinalPrice()
		if err != nil {
			return err
		}

	case constant.Manufactured:
		// manufactured: 12.5% of the item cost + 2% of (item cost + 12.5% of the item cost)
		err := item.CalculateManufcaturedTax()
		if err != nil {
			return err
		}
		err = item.CalculateManufacturedFinalPrice()
		if err != nil {
			return err
		}
	case constant.Imported:
		//imported: 10% import duty on item cost + a surcharge
		err := item.CalculateImportTax()
		if err != nil {
			return err
		}
		err = item.CalculateImportFinalPrice()
		if err != nil {
			return err
		}
		err = item.ApplyImportSurcharge()
		if err != nil {
			return err
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

func (item *Item) SetItemDetails() error {

	fmt.Printf("Item Name: ")
	_, err := fmt.Scanf("%s", &(item.Name))
	if err != nil {
		log.Println("scan for Item Name failed, due to ", err)
		return err
	}

	fmt.Printf("Item Price: ")
	_, err = fmt.Scanf("%g", &(item.Price))
	if err != nil {
		log.Println("scan for Item Price failed, due to ", err)
		return err
	}

	fmt.Printf("Item Quantity: ")
	_, err = fmt.Scanf("%d", &(item.Quantity))
	if err != nil {
		log.Println("scan for Item Quantity failed, due to ", err)
		return err
	}

	fmt.Printf("Item Type: ")
	_, err = fmt.Scanf("%s", &(item.Type))
	if err != nil {
		log.Println(" scan for Item type failed, due to ", err)
		return err
	}

	err = item.ValidateItemDetails()
	if err != nil {
		log.Println(err.Error())
		err = item.SetItemDetails()
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil

}

func (item *Item) ValidateItemDetails() error {
	if len(item.Type) == 0 {
		return custErr.InvalideItmType
	}
	if item.Quantity < 0 {
		return custErr.NegativeQuantErr
	}
	if item.Price < 0 {
		return custErr.NegativePriceErr
	}
	if item.Type != constant.Raw && item.Type != constant.Manufactured && item.Type != constant.Imported {
		return custErr.InvalideItmType
	}
	return nil
}
