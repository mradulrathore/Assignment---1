package services

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strings"

	models "mradulrathore/onboarding-assignments/assignment2/models"
)

var usersDetails = []models.User{}

func AddUserDetails() (ok bool, err error) {
	user, err := GetUserDetails()
	for err != nil {
		fmt.Println(err)
		user, err = GetUserDetails()
	}
	_, err = InsertUserDetails(user)

	if err != nil {
		return false, err
	}
	return true, nil
}

func GetUserDetails() (user models.User, err error) {
	fmt.Printf("Full Name: ")
	_, err = fmt.Scanf("%s", &(user.FullName))
	if err != nil {
		log.Println("scan for user's name failed, due to ", err)
		return models.User{}, err
	}

	fmt.Printf("Age: ")
	_, err = fmt.Scanf("%d", &(user.Age))
	if err != nil {
		log.Println("scan for user's age failed, due to ", err)
		return models.User{}, err
	}

	fmt.Printf("Address: ")
	_, err = fmt.Scanf("%s", &(user.Address))
	if err != nil {
		log.Println("scan for user's address failed, due to ", err)
		return models.User{}, err
	}

	fmt.Printf("Roll No : ")
	_, err = fmt.Scanf("%d", &(user.RollNo))
	if err != nil {
		log.Println(" scan for user's rollno failed, due to ", err)
		return models.User{}, err
	}

	fmt.Printf("Courses : ")
	var coursesEnrol string
	_, err = fmt.Scanf("%s", &coursesEnrol)
	if err != nil {
		log.Println(" scan for course choice failed, due to ", err)
		return models.User{}, err
	}

	_, err = ValidateCourse(coursesEnrol)
	if err != nil {
		return models.User{}, err
	}
	user.CoursesEnrol = models.Course{Name: coursesEnrol}

	ok, err := ValidateUserDetails(user)

	if !ok {
		log.Println(err.Error())
		return models.User{}, err
	}

	return user, nil
}

// insertAt inserts v into s at index i and returns the new slice.
func InsertAt(index int, user models.User) error {
	if index == len(usersDetails) {
		// Insert at end is the easy case.
		usersDetails = append(usersDetails, user)
		return nil
	}

	// Make space for the inserted element by shifting
	// values at the insertion index up one index. The call
	// to append does not allocate memory when cap(data) is
	// greater ​than len(data).
	usersDetails = append(usersDetails[:index+1], usersDetails[index:]...)

	// Insert the new element.
	usersDetails[index] = user

	// Return the updated slice.
	return nil
}

func InsertUserDetails(user models.User) (bool, error) {

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
		return false, err
	}

	return true, nil

}

func ValidateUserDetails(user models.User) (ok bool, err error) {

	if user.Address == "" {
		errorMessage := "address can not be blank"
		log.Println(errorMessage)
		return false, errors.New(errorMessage)
	}
	if user.RollNo == 0 {
		errorMessage := "please provide rollno"
		log.Println(errorMessage)
		return false, errors.New(errorMessage)
	}
	if user.FullName == "" {
		errorMessage := "name can not be blank"
		log.Println(errorMessage)
		return false, errors.New(errorMessage)
	}
	if user.RollNo < 0 {
		errorMessage := "roll no can not be negative"
		log.Println(errorMessage)
		return false, errors.New(errorMessage)
	}
	if user.Age < 0 {
		errorMessage := "age no can not be negative"
		log.Println(errorMessage)
		return false, errors.New(errorMessage)
	}
	return true, nil
}

func DisplayUserDetails() {

}

func DeleteUserDetails() {

}

func SaveUserDetails() {

}

func Exit() {

}
