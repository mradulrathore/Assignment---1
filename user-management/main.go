package main

import (
	"log"
	"mradulrathore/onboarding-assignments/user-management/ports"
)

func main() {
	err := ports.Init()
	if err != nil {
		log.Println(err)
	}
}
