package view

import (
	"fmt"
	"log"

	"github.com/mradulrathore/onboarding-assignments/constant"
	custErr "github.com/mradulrathore/onboarding-assignments/error"
	itm "github.com/mradulrathore/onboarding-assignments/item"
)

func GetUserChoice() (string, error) {

	fmt.Println("Do you want to enter details of any other item (" + constant.Accept + "/" + constant.Deny + ")")
	var moreItems string = constant.Accept
	_, err := fmt.Scanf("%s", &moreItems)
	if err != nil {
		log.Println(err)
		return "", err
	}

	err = ValidateConfirmation(moreItems)

	for err != nil {

		_, err = fmt.Scanf("%s", &moreItems)
		if err != nil {
			log.Println(err)
			return "", err
		}
		err = ValidateConfirmation(moreItems)
	}

	return moreItems, nil
}

// validate whether userChoice is eiter Accept or Deny
func ValidateConfirmation(userChoice string) error {

	if userChoice != constant.Accept && userChoice != constant.Deny {
		log.Println("enter either " + constant.Accept + " or " + constant.Deny)
		return custErr.InvalidUsrChoice
	}

	return nil
}

func Initialize(item itm.Item) {

	var items []itm.Item

	err := item.ValidateItemDetails()
	if err != nil {
		log.Println(err.Error())
		err = item.SetItemDetails()
		if err != nil {
			//logging is already done in SetItemDetails()
			log.Fatal(err)
		}
	}

	err = item.CalculateTaxAndPrice()
	if err != nil {
		log.Fatal(err)
	}

	items = append(items, item)

	// check whether user wants to add more item
	var moreItems string
	moreItems, err = GetUserChoice()
	if err != nil {
		log.Fatal(err)
	}

	// accept items details from user iteratively
	for moreItems == constant.Accept {

		err = item.SetItemDetails()
		if err != nil {
			log.Fatal(err)
		}

		err = item.CalculateTaxAndPrice()
		if err != nil {
			log.Fatal(err)
		}

		items = append(items, item)

		moreItems, err = GetUserChoice()
		if err != nil {
			log.Fatal(err)
		}
	}

	err = itm.GetAllItemDetails(items)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(items)
}
