package repository

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	usr "github.com/mradulrathore/user-management/service/user"
)

const (
	UserExistErr    = "user exist with id:%d"
	UserNotExistErr = "user does not exist with id:%d"
	Name            = "name"
	RollNo          = "rollno"
	Address         = "address"
	Age             = "age"
)

type Repository interface {
	Load(dataFilePath string) error
	Add(user usr.User) error
	List(field string, ASCOrder bool) (users []usr.User, err error)
	Delete(rollno int) error
	Save(users []usr.User) error
	Close() error
}

type repository struct {
	users map[int]usr.User
	file  *os.File
}

func NewRepo() *repository {
	return &repository{}
}

func (r *repository) Load(dataFilePath string) error {
	if err := open(r, dataFilePath); err != nil {
		return err
	}

	r.users = make(map[int]usr.User)

	usrs, err := retrieveData(r)
	if err != nil {
		return err
	}

	for _, user := range usrs {
		r.users[user.RollNo] = user
	}

	return nil
}

func open(r *repository, dataFilePath string) error {
	if r.file != nil {
		return nil
	}

	file, err := os.OpenFile(dataFilePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Println(err)
		return err
	}

	r.file = file

	return nil
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

	users, err := usr.DecodeUsers(dataB)
	if err != nil {
		return []usr.User{}, err
	}

	return users, nil
}

func (r *repository) Add(user usr.User) error {
	if _, exist := r.users[user.RollNo]; exist {
		err := fmt.Errorf(UserExistErr, user.RollNo)
		log.Println(err)
		return err
	}

	r.users[user.RollNo] = user
	return nil
}

func (r *repository) List(field string, ASCOrder bool) ([]usr.User, error) {
	var users []usr.User
	for _, user := range r.users {
		users = append(users, user)
	}

	if ASCOrder {
		sortAscCustom(users, field)
	} else {
		sortDescCustom(users, field)
	}

	return users, nil
}

func sortAscCustom(users []usr.User, field string) {
	sort.SliceStable(users, func(i, j int) bool {
		switch field {
		case Name:
			return (strings.Compare(users[i].Name, users[j].Name) == -1)
		case RollNo:
			return (users[i].RollNo < users[j].RollNo)
		case Address:
			return (strings.Compare(users[i].Address, users[j].Address) == -1)
		case Age:
			return (users[i].Age < users[j].Age)
		default:
			return true
		}
	})
}

func sortDescCustom(users []usr.User, field string) {
	sort.SliceStable(users, func(i, j int) bool {
		switch field {
		case Name:
			return (strings.Compare(users[i].Name, users[j].Name) == 1)
		case RollNo:
			return (users[i].RollNo > users[j].RollNo)
		case Address:
			return (strings.Compare(users[i].Address, users[j].Address) == 1)
		case Age:
			return (users[i].Age > users[j].Age)
		default:
			return true
		}
	})
}

func (r *repository) Delete(rollno int) error {
	if _, exist := r.users[rollno]; !exist {
		err := fmt.Errorf(UserNotExistErr, rollno)
		log.Println(err)
		return err
	}

	delete(r.users, rollno)

	return nil
}

func (r *repository) Save(users []usr.User) error {
	dataB, err := usr.EncodeUsers(users)
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
