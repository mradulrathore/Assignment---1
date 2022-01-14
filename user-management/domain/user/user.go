package user

import (
	"errors"
	"fmt"
	"log"

	cours "github.com/mradulrathore/user-management/domain/course"

	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	Name         string       `json:"fullname"`
	Age          int          `json:"age"`
	Address      string       `json:"address"`
	RollNo       int          `json:"rollno"`
	CoursesEnrol cours.Course `json:"coursesenrol"`
}

func New(name string, age int, address string, rollNo int, courseEnrol []string) (User, error) {

	var user User
	var err error
	user.Name = name
	user.Age = age
	user.Address = address
	user.RollNo = rollNo
	user.CoursesEnrol, err = cours.New(courseEnrol)

	if err != nil {
		return User{}, err
	}

	if err = user.validate(); err != nil {
		log.Println(err)
		return User{}, err
	}

	return user, nil
}

func (user User) validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Name, validation.Required),
		validation.Field(&user.Age, validation.Required, validation.By(checkPositive)),
		validation.Field(&user.Address, validation.Required),
		validation.Field(&user.RollNo, validation.Required, validation.By(checkPositive)),
	)
}

func checkPositive(value interface{}) error {
	val := value.(int)
	if val <= 0 {
		return errors.New("must be positive")
	}
	return nil
}

func (user User) String() string {
	return fmt.Sprintf("[%s, %d, %s,%d,%s]", user.Name, user.Age, user.Address, user.RollNo, user.CoursesEnrol.String())
}
