package repository

import (
	"encoding/json"
	"log"
	"os"

	usr "github.com/mradulrathore/user-management/domain/user"
)

func Open() (*os.File, error) {
	file, err := os.OpenFile("user-data", os.O_RDWR, 0644)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return file, nil
}

func Save(file *os.File, user []usr.User) error {
	dataB, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		return err
	}
	if err = file.Truncate(0); err != nil {
		log.Println(err)
		return err
	}
	_, err = file.Write(dataB)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Get(file *os.File) ([]usr.User, error) {
	var users []usr.User

	fs, err := file.Stat()
	if err != nil {
		log.Println(err)
		return []usr.User{}, err
	}
	len := fs.Size()
	if len == 0 {
		return []usr.User{}, err
	}

	dataB := make([]byte, len)
	_, err = file.Read(dataB)
	if err != nil {
		log.Println(err)
		return []usr.User{}, err
	}

	if err = json.Unmarshal(dataB, &users); err != nil {
		log.Println(err)
		return []usr.User{}, err
	}

	return users, nil
}
