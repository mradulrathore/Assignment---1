package main

import (
	"fmt"

	"github.com/mradulrathore/dependency-graph/view"
)

func main() {
	if err := view.Init(); err != nil {
		fmt.Println(err)
	}
}
