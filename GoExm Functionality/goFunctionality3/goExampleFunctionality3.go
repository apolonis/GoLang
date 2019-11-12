package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go is running")
	fmt.Println("What's ur name: ")

	name := ""

	fmt.Scanln(&name)

	fmt.Println("Your name is " + name)
}
