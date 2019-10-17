package main

import "fmt"

//Interface
type venhicle interface {
	drive()
}

type car struct {
	name string
}
type truck struct {
	name string
}

func (v venhicle) drive() {
	fmt.Println("Car is driving")
}

// func (t truck) drive() {
// fmt.Println("Truck is driving")
// }

func main() {

	c := car{"Mazda"}
	t := truck{"Tamic"}
	fmt.Println(c)
	fmt.Println(t)
	fmt.Println("___________________")

	c.drive()
	t.drive()

}
