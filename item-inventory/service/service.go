package service

import (
	"fmt"
	"log"
)

func Initialize() error {
	db, cleanup, err := Open()
	if err != nil {
		log.Println(err)
	}
	defer cleanup()

	repo := NewRepo(db)

	itemDB := make(chan Item)
	go getItemsFromDB(repo, itemDB)

	var items []Item
	for itm := range itemDB {
		items = append(items, itm)
	}

	itemMemory := make(chan Item)
	go getItemsFromMemory(items, itemMemory)

	for itm := range itemMemory {
		fmt.Println(itm.Invoice())
	}

	return nil
}

func getItemsFromDB(repo *Repository, itemDB chan Item) {
	list, err := repo.GetItems()
	if err != nil {
		log.Println(err)
	}

	for _, item := range list.Items {
		itemDB <- item
	}

	close(itemDB)
}

func getItemsFromMemory(items []Item, itemMemory chan Item) {
	for _, itm := range items {
		itemMemory <- itm
	}

	close(itemMemory)
}
