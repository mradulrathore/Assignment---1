package main

import (
	"log"

	"github.com/mradulrathore/item-inventory/service"
)

func main() {
	if err := service.Initialize(); err != nil {
		log.Println(err)
	}
}
