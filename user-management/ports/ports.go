package ports

import (
	"errors"
	"fmt"
	"log"
	f "mradulrathore/onboarding-assignments/user-management/adapters/file"
	usrApp "mradulrathore/onboarding-assignments/user-management/application/user"
	cours "mradulrathore/onboarding-assignments/user-management/domain/course"
	usr "mradulrathore/onboarding-assignments/user-management/domain/user"
)

func Init() error {

	err := load()
	if err != nil {
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

func load() (err error) {

	file, err := f.Open()
	if err != nil {
		return
	}
	defer file.Close()

	users, err := f.Retrieve(file)
	if err != nil {
		return
	}
	usrApp.Init(users)
	return
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

func getUserChoice() (userChoice string, err error) {
	_, err = fmt.Scanf("%s", &userChoice)
	if err != nil {
		log.Println("scan for user choice failed, due to ", err)
		return
	}
	return
}

func addUser() (err error) {
	name, age, address, rollNo, courseEnrol, err := getUser()
	if err != nil {
		return
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

	usrApp.Insert(user)
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
	fmt.Printf("Enter number of courses you want to enrol (atleast %d) ", cours.MinCousesEnrol)
	var numCourse int
	_, err = fmt.Scanf("%d", &numCourse)
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
		fmt.Println("Enter ", i, "th course")
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
			log.Println(DuplicateCourseErr)
			return DuplicateCourseErr
		} else {
			courseFrequency[course] = 1
		}
	}
	return nil
}

func display() (err error) {
	fmt.Print("Field Name to sort details on (1. Ascending 2.Descending): ")

	var field string
	_, err = fmt.Scanf("%s", &field)
	if err != nil {
		log.Println(err)
		return
	}

	var order int
	_, err = fmt.Scanf("%d", &order)
	if err != nil {
		log.Println(err)
		return
	}

	users := usrApp.GetAll(field, order)
	for _, user := range users {
		fmt.Println(user.String())
	}
	return
}

func deleteByRollNo() (err error) {
	fmt.Print("Enter roll no to delete: ")
	var rollNo int
	_, err = fmt.Scanf("%d", &rollNo)
	if err != nil {
		log.Println(err)
		return
	}
	err = usrApp.DeleteByRollNo(rollNo)
	return
}

func save() (err error) {
	file, err := f.Open()
	if err != nil {
		return
	}
	defer file.Close()

	//saving data in ascending order of name
	users := usrApp.GetAll("name", 1)
	err = f.Save(file, users)
	return
}

func confirmSave() (err error) {
	fmt.Println("Do you want to save the data(y/n)?")
	var userChoice string
	fmt.Scanf("%s", &userChoice)
	err = validateConfirmation(userChoice)
	if err != nil {
		return
	}
	if userChoice == "y" {
		err = save()
	}
	return
}

// validate whether userChoice is eiter Accept or Deny
func validateConfirmation(userChoice string) error {
	if userChoice != Accept && userChoice != Deny {
		log.Println(InvalidUsrChoice)
		return InvalidUsrChoice
	}

	return nil
}
