package main

import "fmt"

//EXAMPLE HOW TO USE SETER IN GO LANG

type user struct {
	name     string
	lastname string
}

func (u *user) setName(name string) {

	u.name = name

}

func main() {

	a := user{"asd", "asd"}

	a.name = "ASD"

	a.setName("NEWNAME")

	fmt.Println(a)

}
