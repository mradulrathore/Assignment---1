package service

import "fmt"

func getItemsFromMemory(itemDB chan Item) {
	item := <-itemDB
	fmt.Println(item.Invoice())
}
