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

	item := make(chan Item)
	go getItem(repo, item)

	for itm := range item {
		fmt.Println(itm.Invoice())
	}

	return nil
}

func getItem(repo *Repository, item chan Item) {
	list, err := repo.GetItems()
	if err != nil {
		log.Println(err)
	}
	for _, itemDb := range list.Items {
		item <- itemDb
	}

	close(item)
}
