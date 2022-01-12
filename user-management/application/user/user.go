package user

import (
	"errors"
	"log"
	usr "mradulrathore/onboarding-assignments/user-management/domain/user"
	"sort"
	"strings"
)

var users = []usr.User{}

func Initialize(usrs []usr.User) {
	users = usrs
}

func Insert(user usr.User) {
	index := searchName(user)
	insertAt(index, user)
}

// return the smallest index i in [0, n) at which user.Name should be inserted to maintain sorted list
// if user.Name matches with already existing element name then return index according to rollno
func searchName(user usr.User) (index int) {
	sort.Search(len(users), func(i int) bool {
		if strings.Compare(users[i].Name, user.Name) == 1 {
			return true
		} else if strings.Compare(users[i].Name, user.Name) == 0 {
			return users[i].RollNo > user.RollNo
		}
		return false
	})

	return
}

func insertAt(index int, user usr.User) {
	if index == len(users) {
		users = append(users, user)
		return
	}

	users = append(users[:index+1], users[index:]...)
	users[index] = user
}

func GetAll(field string, order int) (users []usr.User) {

	if order == 1 {
		sortAscCustom(field)
	} else {
		sortDescCustom(field)
	}

	return
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

func DeleteByRollNo(rollNo int) (err error) {
	index := searchRollNo(rollNo)
	//    return append(slice[:s], slice[s+1:]...)

	if users[index].RollNo != rollNo {
		errMsg := "roll no doesn't exist"
		log.Println(errMsg)
		err = errors.New(errMsg)
		return
	}

	users = append(users[:index], users[index+1:]...)
	return
}

func searchRollNo(rollNo int) (index int) {
	index = sort.Search(len(users), func(i int) bool {
		return users[i].RollNo >= rollNo
	})
	return
}
