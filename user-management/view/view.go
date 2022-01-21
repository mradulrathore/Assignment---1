package view

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	repo "github.com/mradulrathore/user-management/service/repository"

	usr "github.com/mradulrathore/user-management/service/user"
)

const (
	Accept             = "y"
	Deny               = "n"
	DataFilePath       = "user-data.json"
	DuplicateCourseErr = "duplicate course"
	MinCouses          = 4
	TotalCourses       = 6
)

func Init() error {
	repository := repo.NewRepo()
	if err := repository.Load(DataFilePath); err != nil {
		log.Println(err)
	}
	defer repository.Close()

	var moreInput bool = true
	for moreInput {
		showMenu()

		userChoice, err := getUserChoice()
		if err != nil {
			return err
		}
		switch userChoice {
		case "1":
			err := addUser(repository)
			if err != nil {
				fmt.Println(err)
			}
		case "2":
			users, err := getAll(repository)
			if err != nil {
				fmt.Println(err)
			}
			display(users)
		case "3":
			err := deleteByRollNo(repository)
			if err != nil {
				fmt.Println(err)
			}
		case "4":
			if err = save(repository); err != nil {
				fmt.Println(err)
			}
		case "5":
			moreInput = false
			err = confirmSave(repository)
			if err != nil {
				moreInput = true
			} else {
				fmt.Println("exiting")
			}
		default:
			fmt.Println("Invalid choice")
		}
	}

	return nil
}

func showMenu() {
	fmt.Println()
	fmt.Println("-------------------")
	fmt.Println("1. Add user details")
	fmt.Println("2. Display user details")
	fmt.Println("3. Delete user details")
	fmt.Println("4. Save user details")
	fmt.Println("5. Exit")
	fmt.Println("-------------------")
	fmt.Println()
}

func getUserChoice() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	var userChoice string
	if scanner.Scan() {
		userChoice = scanner.Text()
		userChoice = strings.TrimSpace(userChoice)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return "", err
	}
	return userChoice, nil
}

func addUser(repository repo.Repository) error {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Full Name: ")
	var name string
	if scanner.Scan() {
		name = scanner.Text()
		name = strings.TrimSpace(name)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return err
	}

	var age int
	var err error
	fmt.Printf("Age: ")
	if scanner.Scan() {
		age, err = strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Println(err)
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return err
	}

	fmt.Printf("Address: ")
	var address string
	if scanner.Scan() {
		address = scanner.Text()
		address = strings.TrimSpace(address)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return err
	}

	var rollNo int
	fmt.Printf("Roll No : ")
	if scanner.Scan() {
		rollNo, err = strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Println(err)
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	courses, err := getCourse()
	if err != nil {
		return err
	}

	user, err := usr.New(name, age, address, rollNo, courses)
	if err != nil {
		return err
	}

	if err := repository.Add(user); err != nil {
		return err
	}

	fmt.Print("\nuser added successfully\n")

	return nil
}

func getCourse() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	var courses []string
	var err error
	fmt.Printf("Enter number of courses you want to enrol (atleast %d) ", MinCouses)
	var numCourse int
	if scanner.Scan() {
		numCourse, err = strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Println(err)
			return []string{}, err
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return []string{}, err
	}

	if numCourse < MinCouses {
		err := fmt.Errorf("select atleast %d", MinCouses)
		return []string{}, err
	}

	for i := 1; i <= numCourse; i++ {
		fmt.Printf("Enter course - %d: ", i)

		var course string
		if scanner.Scan() {
			course = scanner.Text()
			course = strings.TrimSpace(course)
		}
		if err := scanner.Err(); err != nil {
			log.Println(err)
			return []string{}, err
		}

		courses = append(courses, course)
	}

	if err = checkDuplicateCourse(courses); err != nil {
		return []string{}, err
	}

	return courses, nil
}

func checkDuplicateCourse(courses []string) error {
	courseFrequency := make(map[string]int)
	for _, course := range courses {
		_, exist := courseFrequency[course]
		if exist {
			err := fmt.Errorf("%s", DuplicateCourseErr)
			log.Println(err)
			return err
		} else {
			courseFrequency[course] = 1
		}
	}
	return nil
}

func getAll(repository repo.Repository) ([]usr.User, error) {
	fmt.Print("Field Name to sort details on: ")

	scanner := bufio.NewScanner(os.Stdin)
	var field string
	if scanner.Scan() {
		field = scanner.Text()
		field = strings.TrimSpace(field)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return []usr.User{}, err
	}

	fmt.Print("\n1. Ascending 2.Descending : ")
	var order int
	var err error
	if scanner.Scan() {
		order, err = strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Println(err)
			return []usr.User{}, err
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return []usr.User{}, err
	}

	users, err := repository.GetAll(field, order)
	if err != nil {
		return []usr.User{}, err
	}

	return users, nil
}

func display(users []usr.User) {
	fmt.Print("\n	Name	|	Age	|	Address	|	RollNo	|	Courses	|\n")
	fmt.Println()
	for _, user := range users {
		fmt.Println(user.String())
	}
}

func deleteByRollNo(repository repo.Repository) error {
	fmt.Print("Enter roll no to delete: ")

	scanner := bufio.NewScanner(os.Stdin)
	var rollNo int
	var err error
	if scanner.Scan() {
		rollNo, err = strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Println(err)
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return err
	}

	if err = repository.DeleteByRollNo(rollNo); err != nil {
		return err
	}

	fmt.Print("\nuser deleted successfully\n")

	return nil
}

func save(repository repo.Repository) error {
	//saving data in ascending order of name
	users, err := repository.GetAll("name", 1)
	if err != nil {
		return err
	}

	if err = repository.Save(users); err != nil {
		return err
	}

	fmt.Println("saved successfully")

	return nil
}

func confirmSave(repository repo.Repository) error {
	fmt.Println("Do you want to save the data(y/n)?")

	scanner := bufio.NewScanner(os.Stdin)
	var userChoice string
	if scanner.Scan() {
		userChoice = scanner.Text()
		userChoice = strings.TrimSpace(userChoice)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return err
	}

	if err := validateConfirmation(userChoice); err != nil {
		return err
	}

	if userChoice == "y" {
		if err := save(repository); err != nil {
			return err
		}
	}
	return nil
}

func validateConfirmation(userChoice string) error {
	if userChoice != Accept && userChoice != Deny {
		err := fmt.Errorf("%s", "invalid choice")
		log.Println(err)
		return err
	}

	return nil
}
