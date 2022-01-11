package user

import (
	"fmt"
	"log"
	"sort"
	"strings"

	cours "mradulrathore/onboarding-assignments/assignment2/domain/course"
	usr "mradulrathore/onboarding-assignments/assignment2/domain/user"
	e "mradulrathore/onboarding-assignments/assignment2/error"
	courseService "mradulrathore/onboarding-assignments/assignment2/service/course"
)

var usersDetails = []usr.User{}

func AddUserDetails() (ok bool, err error) {
	user, err := GetUserDetails()
	for err != nil {
		fmt.Println(err)
		user, err = GetUserDetails()
	}
	err = InsertUserDetails(user)

	if err != nil {
		return false, err
	}
	return true, nil
}

func GetUserDetails() (user usr.User, err error) {
	fmt.Printf("Full Name: ")
	_, err = fmt.Scanf("%s", &(user.FullName))
	if err != nil {
		log.Println("scan for user's name failed, due to ", err)
		return usr.User{}, err
	}

	fmt.Printf("Age: ")
	_, err = fmt.Scanf("%d", &(user.Age))
	if err != nil {
		log.Println("scan for user's age failed, due to ", err)
		return usr.User{}, err
	}

	fmt.Printf("Address: ")
	_, err = fmt.Scanf("%s", &(user.Address))
	if err != nil {
		log.Println("scan for user's address failed, due to ", err)
		return usr.User{}, err
	}

	fmt.Printf("Roll No : ")
	_, err = fmt.Scanf("%d", &(user.RollNo))
	if err != nil {
		log.Println(" scan for user's rollno failed, due to ", err)
		return usr.User{}, err
	}

	fmt.Printf("Courses : ")
	var coursesEnrol string
	_, err = fmt.Scanf("%s", &coursesEnrol)
	if err != nil {
		log.Println(" scan for course choice failed, due to ", err)
		return usr.User{}, err
	}

	_, err = courseService.ValidateCourse(coursesEnrol)
	if err != nil {
		return usr.User{}, err
	}
	user.CoursesEnrol = cours.Course{Name: coursesEnrol}

	err = ValidateUserDetails(user)

	if err != nil {
		log.Println(err.Error())
		return usr.User{}, err
	}

	return user, nil
}

func InsertAt(index int, user usr.User) error {
	if index == len(usersDetails) {
		usersDetails = append(usersDetails, user)
		return nil
	}

	usersDetails = append(usersDetails[:index+1], usersDetails[index:]...)

	usersDetails[index] = user

	return nil
}

func InsertUserDetails(user usr.User) error {

	index := sort.Search(len(usersDetails), func(i int) bool {
		if strings.Compare(usersDetails[i].FullName, user.FullName) == 1 {
			return true
		} else if strings.Compare(usersDetails[i].FullName, user.FullName) == 0 {
			return usersDetails[i].RollNo >= user.RollNo
		}
		return false
	})

	err := InsertAt(index, user)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}

func ValidateUserDetails(user usr.User) error {

	if user.Address == "" {
		log.Println(e.EmptyAddr)
		return e.EmptyAddr
	}
	if user.RollNo == 0 {
		log.Println(e.EmptyRollNo)
		return e.EmptyRollNo
	}
	if user.FullName == "" {
		log.Println(e.EmptyFullName)
		return e.EmptyFullName
	}
	if user.RollNo < 0 {
		log.Println(e.NegativeRollNoErr)
		return e.NegativeRollNoErr
	}
	if user.Age < 0 {
		log.Println(e.NegativeAgeErr)
		return e.NegativeAgeErr
	}
	return nil
}

func DisplayUserDetails() {

}

func DeleteUserDetails() {

}

func SaveUserDetails() {

}

func Exit() {

}
