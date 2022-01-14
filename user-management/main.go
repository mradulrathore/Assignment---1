package main

import (
	"log"

	"github.com/mradulrathore/user-management/ports"
)

func main() {
	err := ports.Init()
	if err != nil {
		log.Println(err)
	}
}
