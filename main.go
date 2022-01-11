package main

import (
	"flag"
	"log"

	constant "github.com/mradulrathore/onboarding-assignments/constants"
	itemFile "github.com/mradulrathore/onboarding-assignments/items"
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

	// logging inputted(from command line) item details
	log.Println("item name: ", *name)
	log.Println("price of item: ", *price)
	log.Println("quantity of item: ", *quantity)
	log.Println("type of item: ", *typeItem)

	var itemsDetails []itemFile.Item
	item := itemFile.Item{}

	item.Name = *name
	item.Price = *price
	item.Quantity = *quantity
	item.TypeItem = *typeItem

	// validate input
	ok, err := item.ValidateItemDetails()
	if !ok {
		log.Println(err.Error())
		// take input again from user and validate
		_, err = item.SetItemDetails()
		if err != nil {
			//logging is already done in SetItemDetails()
			log.Fatal(err)
		}
	}

	// calculate tax and final price for the given item
	err = item.CalculateTaxAndPrice()
	if err != nil {
		log.Fatal(err)
	}

	// add item to items list
	itemsDetails = append(itemsDetails, item)

	// check whether user wants to add more item
	var moreItems string
	moreItems, err = itemFile.GetUserChoice()
	if err != nil {
		log.Fatal(err)
	}

	// accept items details from user iteratively
	for moreItems == constant.Accept {

		_, err = item.SetItemDetails()
		if err != nil {
			log.Fatal(err)
		}

		// calculate tax and final price for the given item
		err = item.CalculateTaxAndPrice()
		if err != nil {
			log.Fatal(err)
		}

		// add item to items list
		itemsDetails = append(itemsDetails, item)

		// check whether user wants to enter more item details
		moreItems, err = itemFile.GetUserChoice()
		if err != nil {
			log.Fatal(err)
		}
	}

	// print details of all items
	err = itemFile.GetAllItemDetails(itemsDetails)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(itemsDetails)
}
