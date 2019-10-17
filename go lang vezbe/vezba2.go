package main

import (
	"fmt"
)

//INTERFACE HOW TO USE IT IN GO LANG

type Vehicle interface {
	drive()
}

type Car struct {
	driving string
}

func (c Car) drive() {
	fmt.Println(c.driving)
}

func main() {

	var vehicle Vehicle = Car{"Driving"}

	vehicle.drive()

}
