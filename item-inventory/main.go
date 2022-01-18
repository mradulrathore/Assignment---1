package main

import (
	"log"

	"github.com/mradulrathore/item-inventory/view"
)

func main() {
	if err := view.Initialize(); err != nil {
		log.Println(err)
	}
}
