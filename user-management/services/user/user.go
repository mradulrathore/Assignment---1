package user

import (
	"fmt"
	usr "mradulrathore/onboarding-assignments/user-management/domain/user"
	"sort"
	"strings"
)

var users = []usr.User{}

func Insert(user usr.User) {
	index := sort.Search(len(users), func(i int) bool {
		if strings.Compare(users[i].Name, user.Name) == 1 {
			return true
		} else if strings.Compare(users[i].Name, user.Name) == 0 {
			return users[i].RollNo >= user.RollNo
		}
		return false
	})

	insertAt(index, user)
}

func insertAt(index int, user usr.User) {
	if index == len(users) {
		users = append(users, user)
		return
	}

	users = append(users[:index+1], users[index:]...)
	users[index] = user
}

func Display(field string, order int) {

	if order == 1 {
		sortAscCustom(field)
	} else {
		sortDescCustom(field)
	}

	for _, user := range users {
		fmt.Println(user.String())
	}
}

func sortAscCustom(field string) {
	sort.SliceStable(users, func(i, j int) bool {
		switch field {
		case "name":
			return (strings.Compare(users[i].Name, users[j].Name) == -1)
		case "rollno":
			return (users[i].RollNo < users[j].RollNo)
		case "address":
			return (strings.Compare(users[i].Address, users[j].Address) == -1)
		case "age":
			return (users[i].Age < users[j].Age)
		}
		return true
	})
}

func sortDescCustom(field string) {
	sort.SliceStable(users, func(i, j int) bool {
		switch field {
		case "name":
			return (strings.Compare(users[i].Name, users[j].Name) == 1)
		case "rollno":
			return (users[i].RollNo > users[j].RollNo)
		case "address":
			return (strings.Compare(users[i].Address, users[j].Address) == 1)
		case "age":
			return (users[i].Age > users[j].Age)
		}
		return true
	})
}

//TODO
func Delete() {

}

//TODO
func Save() {

}

//TODO
func Exit() {

}
