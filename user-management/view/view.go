package view

import (
	"fmt"
	"log"

	"github.com/pkg/errors"

	usr "github.com/mradulrathore/user-management/domain/user"
	"github.com/mradulrathore/user-management/repository"
	"github.com/mradulrathore/user-management/service"
)

func Init() error {
	if err := service.LoadData(); err != nil {
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
			if err = add(); err != nil {
				fmt.Println(err)
			}
		case "2":
			if err = display(); err != nil {
				fmt.Println(err)
			}
		case "3":
			if err = deleteByRollNo(); err != nil {
				fmt.Println(err)
			}
		case "4":
			if err = save(); err != nil {
				fmt.Println(err)
			}
		case "5":
			moreInput = false
			if err = confirmSave(); err != nil {
				moreInput = true
			}
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

func getUserChoice() (string, error) {
	var userChoice string
	_, err := fmt.Scanf("%s", &userChoice)
	if err != nil {
		err = errors.Wrap(err, "scan for user choice failed")
		return "", err
	}
	return userChoice, nil
}

func add() error {
	name, age, address, rollNo, courseEnrol, err := getUser()
	if err != nil {
		return err
	}

	user, err := usr.New(name, age, address, rollNo, courseEnrol)
	for err != nil {
		name, age, address, rollNo, courseEnrol, err = getUser()
		if err != nil {
			return err
		}
		user, err = usr.New(name, age, address, rollNo, courseEnrol)
	}

	service.Add(user)

	return nil
}

func getUser() (name string, age int, address string, rollNo int, coursesEnrol []string, err error) {
	fmt.Printf("Full Name: ")
	_, err = fmt.Scanf("%s", &name)
	if err != nil {
		err = errors.Wrap(err, "scan for user's name failed")
		log.Println(err)
		return
	}

	fmt.Printf("Age: ")
	_, err = fmt.Scanf("%d", &age)
	if err != nil {
		err = errors.Wrap(err, "scan for user's age failed")
		log.Println(err)
		return
	}

	fmt.Printf("Address: ")
	_, err = fmt.Scanf("%s", &address)
	if err != nil {
		err = errors.Wrap(err, "scan for user's address failed")
		log.Println(err)
		return
	}

	fmt.Printf("Roll No : ")
	_, err = fmt.Scanf("%d", &rollNo)
	if err != nil {
		err = errors.Wrap(err, "scan for user's rollno failed")
		log.Println(err)
		return
	}
	if err = checkDuplicateRollNo(rollNo); err != nil {
		return
	}

	coursesEnrol, err = getCourse()
	if err != nil {
		return
	}

	return
}

var (
	DuplicateCourseMsg = "duplicate course"
	DuplicateRollNoMsg = "duplicate rollno"
)

func checkDuplicateRollNo(rollno int) error {
	if exist := service.CheckDataExistence(rollno); exist {
		err := fmt.Errorf("%s", DuplicateRollNoMsg)
		log.Println(err)
		return err
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
		err = errors.Wrap(err, "scan for number of course failed")
		log.Println(err)
		return []string{}, err
	}
	if numCourse < MinCousesEnrol {
		err := fmt.Errorf("select atleast %d", MinCousesEnrol)
		return []string{}, err
	}

	for i := 1; i <= numCourse; i++ {
		fmt.Printf("Enter course - %d: ", i)
		var course string
		_, err = fmt.Scanf("%s", &course)
		if err != nil {
			err = errors.Wrapf(err, "scan for course %d failed", i)
			log.Println(err)
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
			err := fmt.Errorf("%s", DuplicateCourseMsg)
			log.Println(err)
			return err
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
		err = errors.Wrap(err, "scan for field name to sort details on failed")
		log.Println(err)
		return err
	}

	var order int
	_, err = fmt.Scanf("%d", &order)
	if err != nil {
		err = errors.Wrap(err, "scan for sorting order failed")
		log.Println(err)
		return err
	}

	users, err := service.GetAll(field, order)
	if err != nil {
		return err
	}
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
		err = errors.Wrap(err, "scan for rollno to delete failed")
		log.Println(err)
		return err
	}
	if err = service.DeleteByRollNo(rollNo); err != nil {
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
	users, err := service.GetAll("name", 1)
	if err != nil {
		return err
	}
	if err = repository.Save(file, users); err != nil {
		return err
	}

	return nil
}

func confirmSave() error {
	fmt.Println("Do you want to save the data(y/n)?")
	var userChoice string
	_, err := fmt.Scanf("%s", &userChoice)
	if err != nil {
		err = errors.Wrap(err, "scan for user choice to save details on exit failed")
		log.Println(err)
		return err
	}

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

func validateConfirmation(userChoice string) error {
	if userChoice != Accept && userChoice != Deny {
		err := fmt.Errorf("%s", "invalid choice")
		log.Println(err)
		return err
	}

	return nil
}
