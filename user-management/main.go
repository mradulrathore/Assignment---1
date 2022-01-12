package main

import (
	"log"
	"mradulrathore/onboarding-assignments/user-management/application"
)

func main() {
	err := application.Init()
	if err != nil {
		log.Println(err)
	}
}
