package service

import (
	"fmt"
	"log"
	"sort"
	"strings"

	usr "github.com/mradulrathore/user-management/domain/user"
	"github.com/mradulrathore/user-management/repository"
)

var users = make(map[int]usr.User)

func LoadData() error {
	file, err := repository.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	usersDisk, err := repository.RetrieveData(file)
	if err != nil {
		return err
	}

	for _, userDisk := range usersDisk {
		users[userDisk.RollNo] = userDisk
	}

	return nil
}

func Add(user usr.User) {
	users[user.RollNo] = user
}

func CheckDataExistence(rollno int) bool {
	_, exists := users[rollno]
	return exists
}

func GetAll(field string, order int) ([]usr.User, error) {

	var usersTemp []usr.User
	for _, user := range users {
		usersTemp = append(usersTemp, user)
	}

	if order == 1 {
		sortAscCustom(usersTemp, field)
	} else {
		sortDescCustom(usersTemp, field)
	}

	return usersTemp, nil
}

func sortAscCustom(usersDisk []usr.User, field string) {
	sort.SliceStable(usersDisk, func(i, j int) bool {
		switch field {
		case "name":
			return (strings.Compare(usersDisk[i].Name, usersDisk[j].Name) == -1)
		case "rollno":
			return (usersDisk[i].RollNo < usersDisk[j].RollNo)
		case "address":
			return (strings.Compare(usersDisk[i].Address, usersDisk[j].Address) == -1)
		case "age":
			return (usersDisk[i].Age < usersDisk[j].Age)
		}
		return true
	})
}

func sortDescCustom(usersDisk []usr.User, field string) {
	sort.SliceStable(usersDisk, func(i, j int) bool {
		switch field {
		case "name":
			return (strings.Compare(usersDisk[i].Name, usersDisk[j].Name) == 1)
		case "rollno":
			return (usersDisk[i].RollNo > usersDisk[j].RollNo)
		case "address":
			return (strings.Compare(usersDisk[i].Address, usersDisk[j].Address) == 1)
		case "age":
			return (usersDisk[i].Age > usersDisk[j].Age)
		}
		return true
	})
}

func DeleteByRollNo(rollno int) error {
	if _, exists := users[rollno]; !exists {
		err := fmt.Errorf("%s", "rollno doesn't exist")
		log.Println(err)
		return err
	}

	delete(users, rollno)
	return nil
}
