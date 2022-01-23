package main

import (
	"fmt"
	"log"

	"github.com/mradulrathore/item-inventory/service"
)

func main() {

	// if err := view.Initialize(); err != nil {
	// 	log.Println(err)
	// }
	db, cleanup, err := service.Open()
	if err != nil {
		log.Println(err)
	}
	defer cleanup()

	repo := service.NewRepo(db)
	list, err := repo.GetItems()
	if err != nil {
		log.Println(err)
	}
	for _, item := range list.Items {
		fmt.Println(item)
	}
}
