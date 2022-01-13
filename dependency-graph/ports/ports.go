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
			err = addDependency()
			if err != nil {
				return err
			}
		case "8":
			err = addNode()
			if err != nil {
				return err
			}
		case "9":
			display()
		case "10":
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
	fmt.Println("9. Display graph")
	fmt.Println("10. Exit")
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

func addDependency() (err error) {

	fmt.Println("Enter ids of nodes")
	var n1 int
	fmt.Scanf("%d", &n1)
	var n2 int
	fmt.Scanf("%d", &n2)

	err = application.AddEdge(n1, n2)
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
	_, exist := application.IdExist(id)
	if exist {
		log.Println(DuplicateIdErr)
		return
	}

	fmt.Printf("Name: ")
	_, err = fmt.Scanf("%s", &name)
	if err != nil {
		log.Println("scan for node's name failed, due to ", err)
		return
	}

	metaData = make(map[string]string)
	err = getAdditionInfo(metaData)

	return
}

func getAdditionInfo(metaData map[string]string) (err error) {
	fmt.Printf("Additional Info (y/n): ")
	var userChoice string
	_, err = fmt.Scanf("%s", &userChoice)
	if err != nil {
		log.Println("scan for user choice for meta data failed, due to ", err)
		return
	}

	if userChoice == "y" {
		var key string
		var value string
		_, err = fmt.Scanf("%s %s", &key, &value)
		if err != nil {
			log.Println("scan for node's meta data failed, due to ", err)
			return
		}
		metaData[key] = value
		err = getAdditionInfo(metaData)
		if err != nil {
			return
		}
	}

	return
}

func display() {
	fmt.Println(application.Display())
}
