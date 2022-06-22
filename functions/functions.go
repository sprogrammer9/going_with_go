package main

import (
	"fmt"
	"strings"
)

func main() { //writing reusable, clean and tidy code

	author := "Oyeyebi Ridwan"
	course := "Going with Go"

	fmt.Println(converter(author, course))

}

func converter(author, course string) (a, c string) {

	author = strings.ToUpper(author)
	course = strings.Title(course)

	return author, course

}
