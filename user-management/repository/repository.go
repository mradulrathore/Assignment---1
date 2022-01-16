package repository

import (
	"log"
	"os"

	usr "github.com/mradulrathore/user-management/domain/user"
)

const dataFilePath = "user-data"

func Open() (*os.File, error) {
	var err error
	file, err := os.OpenFile(dataFilePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return file, err
}

func Save(file *os.File, users []usr.User) error {

	dataB, err := usr.EncodeUser(users)
	if err != nil {
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

func RetrieveData(file *os.File) ([]usr.User, error) {

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

	usersDisk, err := usr.DecodeUser(dataB)
	if err != nil {
		return []usr.User{}, err
	}

	return usersDisk, nil
}
