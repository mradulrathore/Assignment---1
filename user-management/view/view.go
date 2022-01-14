package view

import (
	"errors"
	"fmt"
	"log"

	usr "github.com/mradulrathore/user-management/domain/user"
	"github.com/mradulrathore/user-management/repository"
	usrApp "github.com/mradulrathore/user-management/service"
)

func Init() error {
	if err := load(); err != nil {
		return err
	}

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
			if err != nil {
				return err
			}
		case "2":
			err = display()
			if err != nil {
				return err
			}
		case "3":
			err = deleteByRollNo()
			if err != nil {
				return err
			}
		case "4":
			err = save()
			if err != nil {
				return err
			}
		case "5":
			moreInput = false
			err = confirmSave()
			if err != nil {
				moreInput = true
			}
		default:
			fmt.Println("Invalid choice")
		}
	}
	return nil
}

func load() error {
	file, err := repository.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	users, err := repository.Retrieve(file)
	if err != nil {
		return err
	}

	usrApp.Init(users)

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

func getUserChoice() (string, error) {
	var userChoice string
	_, err := fmt.Scanf("%s", &userChoice)
	if err != nil {
		log.Println("scan for user choice failed, due to ", err)
		return "", err
	}
	return userChoice, nil
}

func addUser() error {
	name, age, address, rollNo, courseEnrol, err := getUser()
	if err != nil {
		return err
	}

	user, err := usr.New(name, age, address, rollNo, courseEnrol)
	for err != nil {
		log.Println(err.Error())
		name, age, address, rollNo, courseEnrol, err = getUser()
		if err != nil {
			return err
		}
		user, err = usr.New(name, age, address, rollNo, courseEnrol)
	}

	usrApp.Insert(user)

	return nil
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
	if err = checkDuplicateRollNo(rollNo); err != nil {
		return
	}

	coursesEnrol, err = getCourse()
	return
}

var (
	InvalidUsrChoice   = errors.New("enter either " + Accept + " or " + Deny)
	DuplicateCourseErr = errors.New("duplicate course")
	DuplicateRollNoErr = errors.New("duplicate rollno")
)

func checkDuplicateRollNo(rollno int) error {
	users := usrApp.GetAll("age", 1)
	index := usrApp.SearchRollNo(rollno)
	if users[index].RollNo == rollno {
		return DuplicateRollNoErr
	}
	return nil
}

const (
	TotalCourses   = 6
	MinCousesEnrol = 4
)

func getCourse() ([]string, error) {
	var coursesEnrol []string
	fmt.Printf("Enter number of courses you want to enrol (atleast %d) ", MinCousesEnrol)
	var numCourse int
	_, err := fmt.Scanf("%d", &numCourse)
	if err != nil {
		log.Println("scan for number of course failed, due to ", err)
		return []string{}, err
	}
	if numCourse < MinCousesEnrol {
		errMsg := fmt.Sprintf("select atleast %d", MinCousesEnrol)
		err = errors.New(errMsg)
		return []string{}, err
	}

	for i := 1; i <= numCourse; i++ {
		fmt.Println("Enter ", i, "th course")
		var course string
		_, err = fmt.Scanf("%s", &course)
		if err != nil {
			log.Printf("scan for course %d failed, due to %g \n ", i, err)
			return []string{}, err
		}
		coursesEnrol = append(coursesEnrol, course)
	}

	if err = checkDuplicateCourse(coursesEnrol); err != nil {
		return []string{}, err
	}

	return coursesEnrol, nil
}

func checkDuplicateCourse(courses []string) error {
	courseFrequency := make(map[string]int)
	for _, course := range courses {
		_, exist := courseFrequency[course]
		if exist {
			log.Println(DuplicateCourseErr)
			return DuplicateCourseErr
		} else {
			courseFrequency[course] = 1
		}
	}
	return nil
}

func display() error {
	fmt.Print("Field Name to sort details on (1. Ascending 2.Descending): ")

	var field string
	_, err := fmt.Scanf("%s", &field)
	if err != nil {
		log.Println(err)
		return err
	}

	var order int
	_, err = fmt.Scanf("%d", &order)
	if err != nil {
		log.Println(err)
		return err
	}

	users := usrApp.GetAll(field, order)
	for _, user := range users {
		fmt.Println(user.String())
	}
	return nil
}

func deleteByRollNo() error {
	fmt.Print("Enter roll no to delete: ")
	var rollNo int
	_, err := fmt.Scanf("%d", &rollNo)
	if err != nil {
		log.Println(err)
		return err
	}
	if err = usrApp.DeleteByRollNo(rollNo); err != nil {
		return err
	}
	return nil
}

func save() error {
	file, err := repository.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	//saving data in ascending order of name
	users := usrApp.GetAll("name", 1)
	if err = repository.Save(file, users); err != nil {
		return err
	}
	return nil
}

func confirmSave() error {
	fmt.Println("Do you want to save the data(y/n)?")
	var userChoice string
	fmt.Scanf("%s", &userChoice)
	if err := validateConfirmation(userChoice); err != nil {
		return err
	}
	if userChoice == "y" {
		if err := save(); err != nil {
			return err
		}
	}
	return nil
}

const (
	Accept = "y"
	Deny   = "n"
)

// validate whether userChoice is eiter Accept or Deny
func validateConfirmation(userChoice string) error {
	if userChoice != Accept && userChoice != Deny {
		log.Println(InvalidUsrChoice)
		return InvalidUsrChoice
	}

	return nil
}
