package view

import (
	"fmt"
	"log"

	"github.com/pkg/errors"

	"github.com/mradulrathore/onboarding-assignments/dependency-graph/service"
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
			if err = getParent(); err != nil {
				fmt.Println(err)
			}
		case "2":
			if err = getChild(); err != nil {
				fmt.Println(err)
			}
		case "3":
			if err = getAncestors(); err != nil {
				fmt.Println(err)
			}
		case "4":
			if err = getDescendants(); err != nil {
				fmt.Println(err)
			}
		case "5":
			if err = deleteDependency(); err != nil {
				fmt.Println(err)
			}
		case "6":
			if err = deleteNode(); err != nil {
				fmt.Println(err)
			}
		case "7":
			if err = addDependency(); err != nil {
				fmt.Println(err)
			}
		case "8":
			if err = addNode(); err != nil {
				fmt.Println(err)
			}
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

func getUserChoice() (string, error) {
	var userChoice string
	_, err := fmt.Scanf("%s", &userChoice)
	if err != nil {
		err := errors.Wrap(err, "scan for user's choice failed")
		log.Println(err)
		return "", err
	}

	return userChoice, nil
}

var (
	IdNotExist = "id:%d doesn't exist"
)

func getParent() error {
	var id int
	fmt.Printf("Enter id: ")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		err = errors.Wrap(err, "scan for node's id (get immediate parent) failed")
		log.Println(err)
		return err
	}

	n, err := service.GetParent(id)
	if err != nil {
		return err
	}

	for _, node := range n {
		fmt.Printf("%v\n", node)
	}

	return nil
}

func getChild() error {
	var id int
	fmt.Printf("Enter id: ")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		err = errors.Wrap(err, "scan for node's id (get immediate child) failed")
		log.Println(err)
		return err
	}

	n, err := service.GetChild(id)
	if err != nil {
		return err
	}

	for _, node := range n {
		fmt.Printf("%v\n", node)
	}

	return nil
}

func getAncestors() error {
	var id int
	fmt.Printf("Enter id: ")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		err = errors.Wrap(err, "scan for node's id (get ancestors) failed")
		log.Println(err)
		return err
	}

	n, err := service.GetAncestors(id)
	if err != nil {
		return err
	}

	for _, node := range n {
		fmt.Printf("%v\n", node)
	}

	return nil
}

func getDescendants() error {
	var id int
	fmt.Printf("Enter id: ")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		err = errors.Wrap(err, "scan for node's id (get descendants) failed")
		log.Println(err)
		return err
	}

	n, err := service.GetDescendants(id)
	if err != nil {
		return err
	}

	for _, node := range n {
		fmt.Printf("%v\n", node)
	}

	return nil
}

func deleteDependency() error {
	fmt.Println("Enter ids of nodes")
	var n1 int
	_, err := fmt.Scanf("%d", &n1)
	if err != nil {
		err := errors.Wrap(err, "scan for node's id-1 failed while adding dependency")
		log.Println(err)
		return err
	}
	var n2 int
	_, err = fmt.Scanf("%d", &n2)
	if err != nil {
		err := errors.Wrap(err, "scan for node's id-2 failed while adding dependency")
		log.Println(err)
		return err
	}

	if err := service.DeleteEdge(n1, n2); err != nil {
		return err
	}

	return nil
}

func deleteNode() error {
	fmt.Println("Enter id of node")
	var id int
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		err := errors.Wrap(err, "scan for node's id failed while deleting")
		log.Println(err)
		return err
	}

	if err := service.DeleteNode(id); err != nil {
		return err
	}

	return nil
}

func addDependency() error {
	fmt.Println("Enter ids of nodes")
	var n1 int
	_, err := fmt.Scanf("%d", &n1)
	if err != nil {
		err := errors.Wrap(err, "scan for node's id-1 failed while adding dependency")
		log.Println(err)
		return err
	}
	var n2 int
	_, err = fmt.Scanf("%d", &n2)
	if err != nil {
		err := errors.Wrap(err, "scan for node's id-2 failed while adding dependency")
		log.Println(err)
		return err
	}

	ancestors, err := service.GetAncestors(n1)
	if err != nil {
		return err
	}
	for _, a := range ancestors {
		if a == n2 {
			err := fmt.Errorf("cyclic dependency")
			log.Println(err)
			return err
		}
	}

	if err = service.AddEdge(n1, n2); err != nil {
		return err
	}

	return nil
}

func addNode() error {
	id, name, metaData, err := getNode()
	if err != nil {
		return err
	}

	if err := service.AddNode(id, name, metaData); err != nil {
		return err
	}

	return nil
}

var (
	DuplicateIdMsg = "duplicate id"
)

func getNode() (id int, name string, metaData map[string]string, err error) {
	fmt.Printf("Id: ")
	_, err = fmt.Scanf("%d", &id)
	if err != nil {
		err = errors.Wrap(err, "scan for node's id failed")
		log.Println(err)
		return
	}
	_, exist := service.CheckIdExist(id)
	if exist {
		err = fmt.Errorf(DuplicateIdMsg)
		log.Println(err)
		return
	}

	fmt.Printf("Name: ")
	_, err = fmt.Scanf("%s", &name)
	if err != nil {
		err = errors.Wrap(err, "scan for node's name failed")
		log.Println(err)
		return
	}

	metaData = make(map[string]string)
	if err = getAdditionInfo(metaData); err != nil {
		err = errors.Wrap(err, "scan for node's metadata failed")
		log.Println(err)
		return
	}

	return
}

func getAdditionInfo(metaData map[string]string) error {
	fmt.Printf("Additional Info (y/n): ")
	var userChoice string
	_, err := fmt.Scanf("%s", &userChoice)
	if err != nil {
		err = errors.Wrap(err, "scan for user's choice for meta data failed")
		log.Println(err)
		return err
	}

	if userChoice == "y" {
		var key string
		var value string
		_, err = fmt.Scanf("%s %s", &key, &value)
		if err != nil {
			err = errors.Wrap(err, "scan for node's meta data failed")
			log.Println(err)
			return err
		}
		metaData[key] = value
		if err = getAdditionInfo(metaData); err != nil {
			return err
		}
	}

	return nil
}
