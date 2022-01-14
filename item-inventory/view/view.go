package view

import (
	"fmt"
	"log"

	"github.com/pkg/errors"

	itm "github.com/mradulrathore/item-inventory/item"
)

const (
	Accept = "y"
	Deny   = "n"
)

func Initialize() error {
	name, price, quantity, typeItem, err := getItem()
	if err != nil {
		return err
	}
	item, err := itm.New(name, price, quantity, typeItem)

	for err != nil {
		log.Println(err.Error())
		name, price, quantity, typeItem, err = getItem()
		if err != nil {
			return err
		}
		item, err = itm.New(name, price, quantity, typeItem)
	}

	fmt.Println(item.Invoice())

	moreItem, err := getUserChoice()
	for err != nil {
		moreItem, err = getUserChoice()
	}

	if moreItem == Accept {
		err = Initialize()
		return err
	}

	return nil
}

func getItem() (name string, price float64, quantity int, typeItem string, err error) {
	fmt.Printf("Item Name: ")
	_, err = fmt.Scanf("%s", &name)
	if err != nil {
		err = errors.Wrap(err, "scan for Item Name failed")
		log.Println(err)
		return
	}

	fmt.Printf("Item Price: ")
	_, err = fmt.Scanf("%f", &price)
	if err != nil {
		err = errors.Wrap(err, "scan for Item Price failed")
		log.Println(err)
		return
	}

	fmt.Printf("Item Quantity: ")
	_, err = fmt.Scanf("%d", &quantity)
	if err != nil {
		err = errors.Wrap(err, "scan for Item Quantity failed")
		log.Println(err)
		return
	}

	fmt.Printf("Item Type: ")
	_, err = fmt.Scanf("%s", &typeItem)
	if err != nil {
		err = errors.Wrap(err, "scan for Item Type failed")
		log.Println(err)
		return
	}

	return
}

func getUserChoice() (string, error) {
	fmt.Println("Do you want to enter details of any other item (" + Accept + "/" + Deny + ")")
	moreItems := Accept
	_, err := fmt.Scanf("%s", &moreItems)
	if err != nil {
		err = errors.Wrap(err, "scan for user choice to enter more item failed")
		log.Println(err)
		return moreItems, err
	}

	if err = validateConfirmation(moreItems); err != nil {
		err = errors.Wrap(err, "user choice validation failed")
		return moreItems, err
	}

	return moreItems, nil
}

func validateConfirmation(userChoice string) error {
	if userChoice != Accept && userChoice != Deny {
		err := fmt.Errorf("%v", "invalid choice")
		log.Println(err)
		return err
	}

	return nil
}
