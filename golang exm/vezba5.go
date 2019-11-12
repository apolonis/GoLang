package main

import "fmt"

//THIS EXAMPLE IS FOR CLASSES THAT HAVE UNIQE ATRIBUTES AND FUNCTIONS, USED WITH MAPS AS WELL

type car struct {
	name     string
	tip      string
	maxspeed int
	on       bool
	speed    int
}

func (c *car) turnOff(i int) {

	c.speed = i
	if c.speed == 0 {

		c.on = false
	}

}

func main() {

	ca := car{"Ferari", "Enzo", 500, true, 50}
	ca2 := car{"Mclaren", "Diablo", 450, false, 0}

	mapa := make(map[string]car)
	mapa["Car1"] = ca
	mapa["Car2"] = ca2

	fmt.Println(mapa)
	fmt.Println("Max speed is:", ca.maxspeed)
	fmt.Println("-----------------")
	fmt.Println("Speed:", ca.speed)
	fmt.Println("-----------------")
	//Here we call function we created above to turn off car
	ca.turnOff(0)
	fmt.Println("Speed is:", ca.speed)
	fmt.Println("-----------------")
	fmt.Println("Car is on?", ca.on)
	fmt.Println("-----------------")
}
