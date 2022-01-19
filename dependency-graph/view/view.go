package view

import (
	"fmt"
	"log"

	graphServ "github.com/mradulrathore/dependency-graph/service/graph"

	nodeServ "github.com/mradulrathore/dependency-graph/service/node"
	"github.com/pkg/errors"
)

func Init() error {

	node := nodeServ.NewNode()
	graph := graphServ.NewGraph(node)

	var moreInput bool = true
	for moreInput {
		showMenu()

		userChoice, err := getUserChoice()
		if err != nil {
			return err
		}
		switch userChoice {
		case "1":
			id, err := getParent()
			if err != nil {
				fmt.Println(err)
			}
			_, err = graph.CheckIdExist(id)
			if err != nil {
				log.Println(err)
				fmt.Println(err)
			}
			nodes, err := graph.GetNodeParent(id)
			if err != nil {
				fmt.Println(err)
			}
			displayNodes(nodes)
		case "2":
			id, err := getChild()
			if err != nil {
				fmt.Println(err)
			}
			_, err = graph.CheckIdExist(id)
			if err != nil {
				log.Println(err)
				fmt.Println(err)
			}
			nodes, err := graph.GetNodeChild(id)
			if err != nil {
				fmt.Println(err)
			}
			displayNodes(nodes)
		case "3":
			id, err := getIDAncestors()
			if err != nil {
				fmt.Println(err)
			}

			ids, err := graph.GetAncestors(id)
			if err != nil {
				fmt.Println(err)
			}

			display(ids)
		case "4":
			id, err := getIDDescendants()
			if err != nil {
				fmt.Println(err)
			}

			ids, err := graph.GetDescendants(id)
			if err != nil {
				fmt.Println(err)
			}

			display(ids)
		case "5":
			id1, id2, err := deleteDependencyIDs()
			if err != nil {
				fmt.Println(err)
			}
			if err := graph.DeleteEdge(id1, id2); err != nil {
				fmt.Println(err)
			}
		case "6":
			id, err := getDeleteNodeID()
			if err != nil {
				fmt.Println(err)
			}
			if err := graph.DeleteNode(id); err != nil {
				fmt.Println(err)
			}
		case "7":
			id1, id2, err := getDependencyIDs()
			if err != nil {
				fmt.Println(err)
			}
			if err = graph.AddEdge(id1, id2); err != nil {
				fmt.Println(err)
			}
		case "8":
			id, name, metaData, err := getNode()
			if err != nil {
				fmt.Println(err)
			}
			if err := graph.AddNode(id, name, metaData); err != nil {
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

func getParent() (int, error) {
	var id int
	fmt.Printf("Enter id: ")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		err = errors.Wrap(err, "scan for node's id (get immediate parent) failed")
		log.Println(err)
		return -1, err
	}

	return id, nil
}

func getChild() (int, error) {
	var id int
	fmt.Printf("Enter id: ")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		err = errors.Wrap(err, "scan for node's id (get immediate child) failed")
		log.Println(err)
		return -1, err
	}

	return id, nil
}

func displayNodes(nodes map[int]nodeServ.Node) {
	for id, _ := range nodes {
		fmt.Println(id)
	}
}

func getIDAncestors() (int, error) {
	var id int
	fmt.Printf("Enter id: ")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		err = errors.Wrap(err, "scan for node's id (get ancestors) failed")
		log.Println(err)
		return -1, err
	}

	return id, nil
}

func display(ids []int) {
	for _, nodeID := range ids {
		fmt.Printf("%v\n", nodeID)
	}
}

func getIDDescendants() (int, error) {
	var id int
	fmt.Printf("Enter id: ")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		err = errors.Wrap(err, "scan for node's id (get descendants) failed")
		log.Println(err)
		return -1, err
	}

	return id, nil
}

func deleteDependencyIDs() (int, int, error) {
	fmt.Println("Enter ids of nodes")
	var n1 int
	_, err := fmt.Scanf("%d", &n1)
	if err != nil {
		err := errors.Wrap(err, "scan for node's id-1 failed while adding dependency")
		log.Println(err)
		return -1, -1, err
	}
	var n2 int
	_, err = fmt.Scanf("%d", &n2)
	if err != nil {
		err := errors.Wrap(err, "scan for node's id-2 failed while adding dependency")
		log.Println(err)
		return -1, -1, err
	}

	return n1, n2, nil
}

func getDeleteNodeID() (int, error) {
	fmt.Println("Enter id of node")
	var id int
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		err := errors.Wrap(err, "scan for node's id failed while deleting")
		log.Println(err)
		return -1, err
	}

	return id, nil
}

func getDependencyIDs() (int, int, error) {
	fmt.Println("Enter ids of nodes")
	var id1 int
	_, err := fmt.Scanf("%d", &id1)
	if err != nil {
		err := errors.Wrap(err, "scan for node's id-1 failed while adding dependency")
		log.Println(err)
		return -1, -1, err
	}
	var id2 int
	_, err = fmt.Scanf("%d", &id2)
	if err != nil {
		err := errors.Wrap(err, "scan for node's id-2 failed while adding dependency")
		log.Println(err)
		return -1, -1, err
	}

	return id1, id2, err
}

func getNode() (int, string, map[string]string, error) {
	var id int
	fmt.Printf("Id: ")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		err = errors.Wrap(err, "scan for node's id failed")
		log.Println(err)
		return -1, "", nil, err
	}

	var name string
	fmt.Printf("Name: ")
	_, err = fmt.Scanf("%s", &name)
	if err != nil {
		err = errors.Wrap(err, "scan for node's name failed")
		log.Println(err)
		return -1, "", nil, err
	}

	metaData := make(map[string]string)
	if err = getAdditionInfo(metaData); err != nil {
		err = errors.Wrap(err, "scan for node's metadata failed")
		log.Println(err)
		return -1, "", nil, err
	}

	return id, name, metaData, nil
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
