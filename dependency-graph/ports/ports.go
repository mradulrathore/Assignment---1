package ports

import (
	"fmt"
	"log"
)

func Init() error {

	var moreInput bool = true
	for moreInput {
		showMenu()
		userChoice, err := getUserChoice()
		if err != nil {
			return err
		}
		switch userChoice {
		case "1":
		case "2":
		case "3":
		case "4":
		case "5":
		case "6":
		case "7":
		case "8":
		case "9":
			moreInput = false
		default:
			fmt.Println("Invalid choice")
		}
	}
	return nil
}

func showMenu() {
	fmt.Println("-------------------")
	fmt.Println("1. Disply the immediate parents of a node")
	fmt.Println("2. Display the immediate children of a node")
	fmt.Println("3. Display the ancestors of a node")
	fmt.Println("4. Display the descendants of a node")
	fmt.Println("5. Delete the dependency")
	fmt.Println("6. Delete the node")
	fmt.Println("7. Add dependency")
	fmt.Println("8. Add node")
	fmt.Println("9. Exit")
	fmt.Println("-------------------")
}

func getUserChoice() (userChoice string, err error) {
	_, err = fmt.Scanf("%s", &userChoice)
	if err != nil {
		log.Println("scan for user choice failed, due to ", err)
		return
	}
	return
}
