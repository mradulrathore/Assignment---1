package file

import (
	"encoding/json"
	"log"
	usr "mradulrathore/onboarding-assignments/user-management/domain/user"
	"os"
)

// If the file doesn't exist, create it, or append to the file
func Open() (file *os.File, err error) {
	file, err = os.OpenFile("user-data", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func Save(file *os.File, user []usr.User) (err error) {
	dataB, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		return
	}
	err = file.Truncate(0)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = file.Write(dataB)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func Retrieve(file *os.File) (users []usr.User, err error) {

	fs, err := file.Stat()
	if err != nil {
		log.Println(err)
		return
	}
	len := fs.Size()
	dataB := make([]byte, len)
	_, err = file.Read(dataB)
	if err != nil {
		log.Println(err)
		return
	}

	err = json.Unmarshal(dataB, &users)
	if err != nil {
		log.Println(err)
	}
	return
}
