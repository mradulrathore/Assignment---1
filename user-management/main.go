package main

import (
	"log"

	"github.com/mradulrathore/user-management/view"
)

func main() {
	if err := view.Init(); err != nil {
		log.Println(err)
	}
}
