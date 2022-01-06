package main

import (
	constant "application/constants"
	itemFile "application/items"
	"flag"
	"fmt"
	"log"
)

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

	var itemsDetails []itemFile.Item

	item := itemFile.Item{}

	item.Name = *name
	item.Price = *price
	item.Quantity = *quantity
	item.TypeItem = *typeItem

	ok, err := item.ValidateItemDetails()
	if !ok {
		fmt.Println(err.Error())
		_, err = item.SetItemDetails()
		if err != nil {
			log.Fatal(err)
		}

	}

	err = item.CalculateTax()
	if err != nil {
		log.Fatal(err)
	}

	itemsDetails = append(itemsDetails, item)

	var moreItems string
	moreItems, err = itemFile.AddMoreItems()
	if err != nil {
		log.Fatal(err)
	}

	for moreItems == constant.Accept {

		_, err = item.SetItemDetails()
		if err != nil {
			log.Fatal(err)
		}

		item.CalculateTax()
		itemsDetails = append(itemsDetails, item)
		moreItems, err = itemFile.AddMoreItems()
		if err != nil {
			log.Fatal(err)
		}
	}

	item.GetItemDetails()
	itemFile.GetAllItemDetails(itemsDetails)
	log.Println(itemsDetails)
}
