package view

import (
	"fmt"
	"log"

	graphServ "github.com/mradulrathore/dependency-graph/service"

	"github.com/pkg/errors"
)

func Init() error {

	graph := graphServ.NewGraph()

	var moreInput bool = true
	for moreInput {
		showMenu()

		userChoice, err := getUserChoice()
		if err != nil {
			return err
		}
		switch userChoice {
		case "1":
			if err := getParent(graph); err != nil {
				fmt.Println(err)
			}
		case "2":
			if err := getChild(graph); err != nil {
				fmt.Println(err)
			}
		case "3":
			if err := getAncestors(graph); err != nil {
				fmt.Println(err)
			}
		case "4":
			if err := getDescendants(graph); err != nil {
				fmt.Println(err)
			}
		case "5":
			if err := deleteDependency(graph); err != nil {
				fmt.Println(err)
			}
		case "6":
			if err := deleteNode(graph); err != nil {
				fmt.Println(err)
			}
		case "7":
			if err := addDependency(graph); err != nil {
				fmt.Println(err)
			}
		case "8":
			if err := addNode(graph); err != nil {
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

func getParent(graph graphServ.Graph) error {
	var id int
	fmt.Printf("Enter id: ")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		err = errors.Wrap(err, "scan for node's id (get immediate parent) failed")
		log.Println(err)
		return err
	}

	nodes, err := graph.GetNodeParent(id)
	if err != nil {
		return err
	}

	fmt.Println(graph.GetNodeDetails(nodes))
	return nil
}

func getChild(graph graphServ.Graph) error {
	var id int
	fmt.Printf("Enter id: ")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		err = errors.Wrap(err, "scan for node's id (get immediate child) failed")
		log.Println(err)
		return err
	}

	nodes, err := graph.GetNodeChild(id)
	if err != nil {
		return err
	}

	fmt.Println(graph.GetNodeDetails(nodes))

	return nil
}

func getAncestors(graph graphServ.Graph) error {
	var id int
	fmt.Printf("Enter id: ")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		err = errors.Wrap(err, "scan for node's id (get ancestors) failed")
		log.Println(err)
		return err
	}

	nodes, err := graph.GetAncestors(id)
	if err != nil {
		return err
	}

	fmt.Println(graph.GetNodeDetails(nodes))
	return nil
}

func getDescendants(graph graphServ.Graph) error {
	var id int
	fmt.Printf("Enter id: ")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		err = errors.Wrap(err, "scan for node's id (get descendants) failed")
		log.Println(err)
		return err
	}

	nodes, err := graph.GetDescendants(id)
	if err != nil {
		return err
	}

	fmt.Println(graph.GetNodeDetails(nodes))

	return nil
}

func deleteDependency(graph graphServ.Graph) error {
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

	if err := graph.DeleteEdge(n1, n2); err != nil {
		return err
	}

	return nil
}

func deleteNode(graph graphServ.Graph) error {
	fmt.Println("Enter id of node")
	var id int
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		err := errors.Wrap(err, "scan for node's id failed while deleting")
		log.Println(err)
		return err
	}

	if err := graph.DeleteNode(id); err != nil {
		return err
	}

	return nil
}

func addDependency(graph graphServ.Graph) error {
	fmt.Println("Enter ids of nodes")
	var id1 int
	_, err := fmt.Scanf("%d", &id1)
	if err != nil {
		err := errors.Wrap(err, "scan for node's id-1 failed while adding dependency")
		log.Println(err)
		return err
	}
	var id2 int
	_, err = fmt.Scanf("%d", &id2)
	if err != nil {
		err := errors.Wrap(err, "scan for node's id-2 failed while adding dependency")
		log.Println(err)
		return err
	}

	if err = graph.AddEdge(id1, id2); err != nil {
		return err
	}

	return err
}

func addNode(graph graphServ.Graph) error {
	var id int
	fmt.Printf("Id: ")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		err = errors.Wrap(err, "scan for node's id failed")
		log.Println(err)
		return err
	}

	var name string
	fmt.Printf("Name: ")
	_, err = fmt.Scanf("%s", &name)
	if err != nil {
		err = errors.Wrap(err, "scan for node's name failed")
		log.Println(err)
		return err
	}

	metaData := make(map[string]string)
	if err = getAdditionInfo(metaData); err != nil {
		err = errors.Wrap(err, "scan for node's metadata failed")
		log.Println(err)
		return err
	}

	if err := graph.AddNode(id, name, metaData); err != nil {
		return err
	}

	return nil
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
