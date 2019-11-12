package main

import (
	"fmt"
	"strconv"
	"time"
)

type venhicle interface {
	drive(drivingTime int)
}

type car struct {
	name string
}
type truck struct {
	name string
}

type motocycle struct {
	name string
}

func (c car) drive(drivingTime int) {
	var i int = 0

	var k string //Parsing from inta to string

	for i <= drivingTime {
		k = strconv.Itoa(i)
		fmt.Println("Car " + c.name + " Is driving " + k + " second")
		time.Sleep(time.Second)
		i++
	}

}
func (c truck) drive(drivingTime int) {
	var i int = 0

	var k string //Parsing from inta to string

	for i <= drivingTime {
		k = strconv.Itoa(i)
		fmt.Println("Truck " + c.name + " Is driving " + k + " second")
		time.Sleep(time.Second)
		i++
	}

}
func (c motocycle) drive(drivingTime int) {
	var i int = 0

	var k string //Parsing from inta to string

	for i <= drivingTime {
		k = strconv.Itoa(i)
		fmt.Println("Motocycle " + c.name + " Is driving " + k + " second")
		time.Sleep(time.Second)
		i++
	}

}

func main() {
	fmt.Println("Go is running")
	c := car{"Mazda"}
	t := truck{"Tamic"}
	m := motocycle{"Honda"}
	go c.drive(5)
	go t.drive(5)
	go m.drive(5)

	fmt.Scanln()
}
