package main

import (
	"fmt"
	"log"

	"mradulrathore/onboarding-assignments/assignment2/models"
	"mradulrathore/onboarding-assignments/assignment2/services"
)

func GetUserChoice() (userChoice string, err error) {

	_, err = fmt.Scanf("%s", &userChoice)
	if err != nil {
		log.Println("scan for user choice failed, due to ", err)
		return userChoice, err
	}
	return userChoice, nil
}

func showMenu() error {
	fmt.Println("-------------------")
	fmt.Println("1. Add user details")
	fmt.Println("2. Display user details")
	fmt.Println("3. Delete user details")
	fmt.Println("4. Save user details")
	fmt.Println("5. Exit")
	fmt.Println("-------------------")

	return nil
}

func Init() {

	var moreInput bool = true
	for moreInput {
		err := showMenu()
		if err != nil {
			log.Fatal(err)
		}
		userChoice, err := GetUserChoice()
		if err != nil {
			log.Fatal(err)
		}

		switch userChoice {
		case "1":
			_, err = services.AddUserDetails()
			if err != nil {
				log.Fatal(err)
			}
		case "2":
		case "3":
		case "4":
		case "5":
			moreInput = false
		default:
			fmt.Println("Invalid choice")
		}
	}

}

func main() {

	Init()

	user := models.User{
		FullName: "Mradul Rathore",
		Age:      20,
		Address:  "Indore, M.P.",
		RollNo:   43,
	}

	userB, _ := user.EncodeUser()

	user1, _ := models.DecodeUser(userB)
	fmt.Print(user1)

}
