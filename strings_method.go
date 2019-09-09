package main

import (
	"fmt"
	"strings"
)

func main() {

	str0 := "Welcome To India"

	//split string

	str1 := strings.Split(str0, "To")

	fmt.Println(str1)

	//Contains
	str2 := strings.Contains(str0, "To")
	fmt.Println(str2)

	//Index of
	str3 := strings.Index(str0, "To")

	fmt.Println(str3)

	//Index any

	str4 := strings.IndexAny(str0, "T")
	fmt.Println(str4)

	//count character

	str5 := strings.Count(str0, "e")
	fmt.Println(str5)

}
