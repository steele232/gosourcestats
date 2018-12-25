package main

import (
	"fmt"
	"io/ioutil"
)

func happy() {
	fmt.Println("Heyo!")

	// filebytes, _ := ioutil.ReadFile("en_main.go")
	duh, _ := ioutil.ReadDir("./")
	fmt.Println(duh)
}
