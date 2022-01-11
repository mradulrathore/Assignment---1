package view

import (
	"fmt"
	"log"

	itm "github.com/mradulrathore/onboarding-assignments/item"
)

func getItem() (name string, price float64, quantity int, typeItem string, err error) {
	fmt.Printf("Item Name: ")
	_, err = fmt.Scanf("%s", &name)
	if err != nil {
		log.Println("scan for Item Name failed, due to ", err)
		return
	}

	fmt.Printf("Item Price: ")
	_, err = fmt.Scanf("%g", &price)
	if err != nil {
		log.Println("scan for Item Price failed, due to ", err)
		return
	}

	fmt.Printf("Item Quantity: ")
	_, err = fmt.Scanf("%d", &quantity)
	if err != nil {
		log.Println("scan for Item Quantity failed, due to ", err)
		return
	}

	fmt.Printf("Item Type: ")
	_, err = fmt.Scanf("%s", &typeItem)
	if err != nil {
		log.Println(" scan for Item type failed, due to ", err)
		return
	}

	return
}

func Initialize() {
	name, price, quantity, typeItem, err := getItem()
	if err != nil {
		log.Fatal(err)
	}
	item, err := itm.New(name, price, quantity, typeItem)

	for err != nil {
		log.Println(err.Error())
		name, price, quantity, typeItem, err = getItem()
		if err != nil {
			log.Fatal(err)
		}
		item, err = itm.New(name, price, quantity, typeItem)
	}

	fmt.Println(item.String())

	// check whether user wants to add more item
	moreItem, err := getUserChoice()
	for err != nil {
		moreItem, err = getUserChoice()
	}

	// accept items details from user iteratively
	if moreItem == Accept {
		Initialize()
	}
}

func getUserChoice() (moreItem string, err error) {
	fmt.Println("Do you want to enter details of any other item (" + Accept + "/" + Deny + ")")
	var moreItems string = Accept
	_, err = fmt.Scanf("%s", &moreItems)
	if err != nil {
		log.Println(err)
		return
	}

	err = validateConfirmation(moreItems)

	return
}

// validate whether userChoice is eiter Accept or Deny
func validateConfirmation(userChoice string) error {
	if userChoice != Accept && userChoice != Deny {
		log.Println(InvalidUsrChoice)
		return InvalidUsrChoice
	}

	return nil
}
