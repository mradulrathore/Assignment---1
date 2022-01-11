package models

import (
	"encoding/json"
	"log"
)

type User struct {
	FullName     string `json:"fullname"`
	Age          int    `json:"age"`
	Address      string `json:"address"`
	RollNo       int    `json:"rollno"`
	CoursesEnrol Course `json:"coursesenrol"`
}

func (user *User) EncodeUser() (userB []byte, err error) {

	userB, err = json.Marshal(user)

	if err != nil {
		log.Println(err)
		return userB, err
	}

	return userB, nil
}

func DecodeUser(userB []byte) (user User, err error) {

	if err := json.Unmarshal(userB, &user); err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil

}
