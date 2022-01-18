package application

import (
	"fmt"

	repo "github.com/mradulrathore/user-management/service/repository"

	usr "github.com/mradulrathore/user-management/service/user"
)

var (
	DuplicateCourseErr = "duplicate course"
	DuplicateRollNoErr = "duplicate rollno"
)

type Application interface {
	Add(user usr.User) error
	GetAll(field string, order int) ([]usr.User, error)
	DeleteByRollNo(rollno int) error
	Save(users []usr.User) error
	ConfirmSave(userChoice string) error
}

type application struct {
	repository *repo.Repository
}

func New(repo *repo.Repository) *application {
	return &application{
		repository: repo,
	}
}

func (app *application) Add(user usr.User) error {
	if err := app.repository.Add(user); err != nil {
		return err
	}

	fmt.Print("\nuser added successfully\n")
	return nil
}

func (app *application) GetAll(field string, order int) ([]usr.User, error) {
	users, err := app.repository.GetAll(field, order)
	if err != nil {
		return []usr.User{}, err
	}

	return users, nil
}

func (app *application) DeleteByRollNo(rollNo int) error {
	if err := app.repository.DeleteByRollNo(rollNo); err != nil {
		return err
	}

	fmt.Print("\nuser deleted successfully\n")
	return nil
}

func (app *application) Save() error {
	//saving data in ascending order of name
	users, err := app.repository.GetAll("name", 1)
	if err != nil {
		return err
	}
	if err = app.repository.Save(users); err != nil {
		return err
	}

	fmt.Println("saved successfully")
	return nil
}

func (app *application) ConfirmSave(userChoice string) error {
	if userChoice == "y" {
		if err := app.Save(); err != nil {
			return err
		}
	}
	fmt.Println("exiting")
	return nil
}
