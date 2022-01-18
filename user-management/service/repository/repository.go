package repository

//go:generate mockgen -source=repository.go -destination=repository_mock.go -package=repository
import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	usr "github.com/mradulrathore/user-management/service/user"
)

const (
	dataFilePath    = "user-data.json"
	UserExistErr    = "user exist with id:%d"
	UserNotExistErr = "user does not exist with id:%d"
)

type repository struct {
	users map[int]usr.User
	file  *os.File
}

type RepositoryI interface {
	Load() error
	Add(usr.User) error
	CheckDataExistence(int) bool
	GetAll(string, int) ([]usr.User, error)
	DeleteByRollNo(int) error
	Save([]usr.User) error
	Close() error
}

func NewRepo() *repository {
	return &repository{}
}

func (r *repository) Load() error {
	if err := open(r); err != nil {
		return err
	}
	r.users = make(map[int]usr.User)

	usersTemp, err := retrieveData(r)
	if err != nil {
		return err
	}

	for _, user := range usersTemp {
		r.users[user.RollNo] = user
	}

	return nil
}

func open(r *repository) error {
	if r.file != nil {
		return nil
	}
	var err error
	r.file, err = os.OpenFile(dataFilePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}

func retrieveData(r *repository) ([]usr.User, error) {
	fs, err := r.file.Stat()
	if err != nil {
		log.Println(err)
		return []usr.User{}, err
	}

	len := fs.Size()
	if len == 0 {
		return []usr.User{}, err
	}
	dataB := make([]byte, len)
	_, err = r.file.Read(dataB)
	if err != nil {
		log.Println(err)
		return []usr.User{}, err
	}

	usersDisk, err := usr.DecodeUser(dataB)
	if err != nil {
		return []usr.User{}, err
	}

	return usersDisk, nil
}

func (r *repository) Add(user usr.User) error {
	if exist := r.CheckDataExistence(user.RollNo); exist {
		err := fmt.Errorf(UserExistErr, user.RollNo)
		log.Println(err)
		return err
	}

	r.users[user.RollNo] = user
	return nil
}

func (r *repository) CheckDataExistence(rollno int) bool {
	_, exists := r.users[rollno]
	return exists
}

func (r *repository) GetAll(field string, order int) ([]usr.User, error) {
	var usersTemp []usr.User
	for _, user := range r.users {
		usersTemp = append(usersTemp, user)
	}

	if order == 1 {
		sortAscCustom(usersTemp, field)
	} else {
		sortDescCustom(usersTemp, field)
	}

	return usersTemp, nil
}

func sortAscCustom(usersDisk []usr.User, field string) {
	sort.SliceStable(usersDisk, func(i, j int) bool {
		switch field {
		case "name":
			return (strings.Compare(usersDisk[i].Name, usersDisk[j].Name) == -1)
		case "rollno":
			return (usersDisk[i].RollNo < usersDisk[j].RollNo)
		case "address":
			return (strings.Compare(usersDisk[i].Address, usersDisk[j].Address) == -1)
		case "age":
			return (usersDisk[i].Age < usersDisk[j].Age)
		}
		return true
	})
}

func sortDescCustom(usersDisk []usr.User, field string) {
	sort.SliceStable(usersDisk, func(i, j int) bool {
		switch field {
		case "name":
			return (strings.Compare(usersDisk[i].Name, usersDisk[j].Name) == 1)
		case "rollno":
			return (usersDisk[i].RollNo > usersDisk[j].RollNo)
		case "address":
			return (strings.Compare(usersDisk[i].Address, usersDisk[j].Address) == 1)
		case "age":
			return (usersDisk[i].Age > usersDisk[j].Age)
		}
		return true
	})
}

func (r *repository) DeleteByRollNo(rollno int) error {
	if exist := r.CheckDataExistence(rollno); !exist {
		err := fmt.Errorf(UserNotExistErr, rollno)
		log.Println(err)
		return err
	}

	delete(r.users, rollno)
	return nil
}

func (r *repository) Save(users []usr.User) error {
	dataB, err := usr.EncodeUser(users)
	if err != nil {
		log.Println(err)
		return err
	}
	if err = r.file.Truncate(0); err != nil {
		return err
	}
	_, err = r.file.Seek(0, 0)
	if err != nil {
		return err
	}
	_, err = r.file.Write(dataB)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *repository) Close() error {
	return r.file.Close()
}
