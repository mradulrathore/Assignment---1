package user

import (
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

	InsertAt(index, user)
}

func InsertAt(index int, user usr.User) {
	if index == len(users) {
		users = append(users, user)
		return
	}

	users = append(users[:index+1], users[index:]...)
	users[index] = user
}

//TODO
func Display() {

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
