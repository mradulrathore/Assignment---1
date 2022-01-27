package service

import (
	"log"
	"time"

	"github.com/mradulrathore/item-inventory/config"
)

const (
	GoroutineCount = 10
)

var (
	items []Item
)

func Initialize() error {
	config := config.LoadAppConfig()

	db, cleanup, err := Open(config)
	if err != nil {
		log.Println(err)
	}
	defer cleanup()

	repo := NewRepo(db)

	list := getItemsFromDB(repo)

	itemDB := make(chan Item, 1)
	go Produce(list, itemDB)

	for i := 0; i < GoroutineCount; i++ {
		go getItemsFromMemory(itemDB)
	}

	time.Sleep(time.Second)

	return nil
}
