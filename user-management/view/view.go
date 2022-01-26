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
		return err
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
			if err := addUser(repository); err != nil {
				fmt.Println(err)
			}
		case "2":
			users, err := getAll(repository)
			if err != nil {
				fmt.Println(err)
			}
			display(users)
		case "3":
			if err := deleteByRollNo(repository); err != nil {
				fmt.Println(err)
			}
		case "4":
			if err := save(repository); err != nil {
				fmt.Println(err)
			}
		case "5":
			moreInput = false
			if err := confirmSave(repository); err != nil {
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

	fmt.Printf("Age: ")

	var age int

	if scanner.Scan() {
		a, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Println(err)
			return err
		}
		age = a
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

	fmt.Printf("Roll No : ")

	var rollNo int

	if scanner.Scan() {
		r, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Println(err)
			return err
		}
		rollNo = r
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

	fmt.Printf("Enter number of courses you want to enrol (atleast %d) ", MinCouses)

	var courses []string
	var numCourse int

	if scanner.Scan() {
		n, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Println(err)
			return []string{}, err
		}
		numCourse = n
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

	if err := checkDuplicateCourse(courses); err != nil {
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
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Field Name to sort details on: ")

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

	var order string

	if scanner.Scan() {
		order = scanner.Text()
		order = strings.TrimSpace(order)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return []usr.User{}, err
	}

	var ASCOrder bool = true

	if order == "2" {
		ASCOrder = false
	}

	users, err := repository.List(field, ASCOrder)
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
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter roll no to delete: ")

	var rollNo int

	if scanner.Scan() {
		r, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Println(err)
			return err
		}
		rollNo = r
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return err
	}

	if err := repository.Delete(rollNo); err != nil {
		return err
	}

	fmt.Print("\nuser deleted successfully\n")

	return nil
}

func save(repository repo.Repository) error {
	//saving data in ascending order of name
	users, err := repository.List("name", true)
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
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Do you want to save the data(" + Accept + "/" + Deny + ")?")

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

	if userChoice == Accept {
		if err := save(repository); err != nil {
			return err
		}
	}
	return nil
}

func validateConfirmation(userChoice string) error {
	if userChoice != Accept && userChoice != Deny {
		err := fmt.Errorf("%s: %s", "invalid choice", userChoice)
		log.Println(err)
		return err
	}

	return nil
}
