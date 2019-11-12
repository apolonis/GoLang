package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	_  = iota //ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println("Time is: ", time.Now())
	fmt.Println("-----------------")

	var m int
	m = 13
	var j int
	j = 9

	mj := 17

	fmt.Println(m, j, mj)
	fmt.Println("-----------------")

	var i int = 42

	var k string //Parsing from inta to string
	k = strconv.Itoa(i)
	fmt.Printf("%v,%T\n", k, k)
	fmt.Println("-----------------")

	//Matrix
	var matrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}

	fmt.Println(matrix)
	fmt.Println("-----------------")
	fmt.Println(matrix[0])
	fmt.Println(matrix[1])
	fmt.Println(matrix[2])
	fmt.Println("-----------------")

	fmt.Println("Welcome!")
	fmt.Println("-----------------")
	r1 := Role{"Admin"}
	r2 := Role{"User"}

}
