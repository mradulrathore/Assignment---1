package main

import (
	"flag"
	"log"

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

	log.Printf("item name: %s \n price of item: %g \n quantity of item: %d \n type of item: %s", *name, *price, *quantity, *typeItem)

	view.Initialize(*name, *price, *quantity, *typeItem)
}
