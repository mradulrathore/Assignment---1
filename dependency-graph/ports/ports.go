package ports

import (
	"fmt"
	"log"

	"github.com/mradulrathore/onboarding-assignments/dependency-graph/application"
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
			addNode()
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

func addNode() (err error) {

	id, name, metaData, err := getNode()
	if err != nil {
		return
	}
	application.AddNode(id, name, metaData)
	return
}

func getNode() (id int, name string, metaData map[string]string, err error) {
	fmt.Printf("Id: ")
	_, err = fmt.Scanf("%d", &id)
	if err != nil {
		log.Println("scan for node's id failed, due to ", err)
		return
	}

	err = checkDuplicateId(id)
	if err != nil {
		return
	}

	fmt.Printf("Name: ")
	_, err = fmt.Scanf("%d", &name)
	if err != nil {
		log.Println("scan for node's name failed, due to ", err)
		return
	}

	metaData = make(map[string]string)
	fmt.Printf("Additional Info (q to stop): ")
	var key string
	var value string
	for key != "q" {
		_, err = fmt.Scanf("%s %s", &key, &value)
		if err != nil {
			log.Println("scan for node's meta data failed, due to ", err)
			return
		}
		metaData[key] = value
	}

	return
}

func checkDuplicateId(id int) (err error) {

	return
}
