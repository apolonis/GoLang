package main

import (
	"fmt"
	"time"
)

//EXAMPLE FOR CONCURRENCY IN GO
func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
func asd(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
func www(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
func eee(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
func rrr(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
func ttt(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
func yyy(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
func main() {

	fmt.Println("Concurrency with GoRoutines")

	go compute(5) // when u call function with 'compute()' and 'go compute()' the difference is
	// -that 'go compute()' is twice faster than 'compute()'
	//-twice faster if u have 2 function, 4 time faster for 4 function and so on
	go compute(5)

	go asd(5)
	go www(5)
	go eee(5)
	go rrr(5)
	go ttt(5)
	go yyy(5)

	fmt.Scanln()
}
