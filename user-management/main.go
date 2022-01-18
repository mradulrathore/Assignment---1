package main

import (
	"fmt"

	"github.com/mradulrathore/user-management/view"
)

func main() {
	if err := view.Init(); err != nil {
		fmt.Println(err)
	}
}
