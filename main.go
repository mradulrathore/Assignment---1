package main

import (
	"flag"
	"log"

	"github.com/mradulrathore/onboarding-assignments/item"
	"github.com/mradulrathore/onboarding-assignments/view"
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

	item := item.Item{}
	item.Name = *name
	item.Price = *price
	item.Quantity = *quantity
	item.Type = *typeItem

	view.Initialize(item)
}
