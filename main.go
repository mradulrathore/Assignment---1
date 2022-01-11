package main

import (
	"flag"
	"log"

	"github.com/mradulrathore/onboarding-assignments/constant"
	itm "github.com/mradulrathore/onboarding-assignments/item"
)

//map command line input (-name, -price, -quantity, -type) to variables
var (
	name     = flag.String("name", "", "item name")
	price    = flag.Float64("price", 0, "price of item")
	quantity = flag.Int("quantity", 0, "quantity of item")
	typeItem = flag.String("type", "", "type of item")
)

func main() {

	flag.Parse()

	log.Println("item name: ", *name)
	log.Println("price of item: ", *price)
	log.Println("quantity of item: ", *quantity)
	log.Println("type of item: ", *typeItem)

	var items []itm.Item

	item := itm.Item{}
	item.Name = *name
	item.Price = *price
	item.Quantity = *quantity
	item.Type = *typeItem

	ok, err := item.ValidateItemDetails()
	if !ok {
		log.Println(err.Error())
		_, err = item.SetItemDetails()
		if err != nil {
			//logging is already done in SetItemDetails()
			log.Fatal(err)
		}
	}

	err = item.CalculateTaxAndPrice()
	if err != nil {
		log.Fatal(err)
	}

	items = append(items, item)

	// check whether user wants to add more item
	var moreItems string
	moreItems, err = itm.GetUserChoice()
	if err != nil {
		log.Fatal(err)
	}

	// accept items details from user iteratively
	for moreItems == constant.Accept {

		_, err = item.SetItemDetails()
		if err != nil {
			log.Fatal(err)
		}

		err = item.CalculateTaxAndPrice()
		if err != nil {
			log.Fatal(err)
		}

		items = append(items, item)

		moreItems, err = itm.GetUserChoice()
		if err != nil {
			log.Fatal(err)
		}
	}

	err = itm.GetAllItemDetails(items)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(items)
}
