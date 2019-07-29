package main

import (
	"fmt"
	"github.com/gobuffalo/packr/v2"

)

func main() {

	box := packr.New("stuff to copy", "./assets")

	list := box.List()
	for _, file := range list {
		fmt.Println("Found item: ", file)
	}
}
