package application

import (
	"errors"
	"fmt"
	"log"
	cours "mradulrathore/onboarding-assignments/user-management/domain/course"
	usr "mradulrathore/onboarding-assignments/user-management/domain/user"
	usrServ "mradulrathore/onboarding-assignments/user-management/services/user"
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
			err = addUser()
			return err
		case "2":
		case "3":
		case "4":
		case "5":
			moreInput = false
		default:
			fmt.Println("Invalid choice")
		}
	}
	return nil
}

func showMenu() {
	fmt.Println("-------------------")
	fmt.Println("1. Add user details")
	fmt.Println("2. Display user details")
	fmt.Println("3. Delete user details")
	fmt.Println("4. Save user details")
	fmt.Println("5. Exit")
	fmt.Println("-------------------")
}

func addUser() (err error) {

	name, age, address, rollNo, courseEnrol, err := getUser()
	if err != nil {
		return err
	}
	user, err := usr.New(name, age, address, rollNo, courseEnrol)

	for err != nil {
		log.Println(err.Error())
		name, age, address, rollNo, courseEnrol, err = getUser()
		if err != nil {
			return
		}
		user, err = usr.New(name, age, address, rollNo, courseEnrol)
	}

	usrServ.Insert(user)

	return
}

func getUserChoice() (userChoice string, err error) {

	_, err = fmt.Scanf("%s", &userChoice)
	if err != nil {
		log.Println("scan for user choice failed, due to ", err)
		return
	}
	return
}

func getUser() (name string, age int, address string, rollNo int, coursesEnrol []string, err error) {
	fmt.Printf("Full Name: ")
	_, err = fmt.Scanf("%s", &name)
	if err != nil {
		log.Println("scan for user's name failed, due to ", err)
		return
	}

	fmt.Printf("Age: ")
	_, err = fmt.Scanf("%d", &age)
	if err != nil {
		log.Println("scan for user's age failed, due to ", err)
		return
	}

	fmt.Printf("Address: ")
	_, err = fmt.Scanf("%s", &address)
	if err != nil {
		log.Println("scan for user's address failed, due to ", err)
		return
	}

	fmt.Printf("Roll No : ")
	_, err = fmt.Scanf("%d", &rollNo)
	if err != nil {
		log.Println(" scan for user's rollno failed, due to ", err)
		return
	}

	coursesEnrol, err = getCourse()

	return
}

func getCourse() (coursesEnrol []string, err error) {

	fmt.Printf("Enter number of courses you want to enrol (atleast %d)", cours.MinCousesEnrol)
	var numCourse int
	_, err = fmt.Scanf("%s", &numCourse)
	if err != nil {
		log.Println("scan for number of course failed, due to ", err)
		return
	}
	if numCourse < cours.MinCousesEnrol {
		errMsg := fmt.Sprintf("select atleast %d", cours.MinCousesEnrol)
		err = errors.New(errMsg)
		return
	}

	for i := 1; i <= numCourse; i++ {
		fmt.Println("Enter ", i, " course")
		var course string
		_, err = fmt.Scanf("%s", &course)
		if err != nil {
			log.Printf("scan for course %d failed, due to %g \n ", i, err)
			return
		}
		coursesEnrol = append(coursesEnrol, course)
	}

	err = checkDuplicateCourse(coursesEnrol)

	return
}

func checkDuplicateCourse(courses []string) error {

	courseFrequency := make(map[string]int)

	for _, course := range courses {

		_, exist := courseFrequency[course]

		if exist {
			return errors.New("duplicate course")
		} else {
			courseFrequency[course] = 1
		}
	}
	return nil
}
