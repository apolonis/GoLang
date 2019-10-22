package main

import (
	"bufio"
	"fmt"
	"os"
)

//Simple calculator with scanner
func calculator() (d float32) {
	var c float32
	var a float32
	var b float32

	fmt.Println("Please enter first number")
	fmt.Println("Calculator is running")
	fmt.Scanln(&a)

	fmt.Println("Please enter the operation")
	reader := bufio.NewReader(os.Stdin)
	x, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Please enter 2nd number")
	fmt.Scanln(&b)

	if x == '+' {
		c = a + b
	} else if x == '-' {
		c = a - b
	} else if x == '*' {
		c = a * b
	} else if x == '/' {
		c = a / b
	} else {
		fmt.Println("Please use +,-,* or / operations!")
	}
	fmt.Print("= ")
	return c
}

func main() {
	fmt.Println("Go is running")

	fmt.Println("--------")

	var q string

	i := 0

	for i < 100 {

		fmt.Println(calculator())
		i++
		fmt.Println("type 'quit' if u want to quit/'enter' to continue")
		fmt.Scanln(&q)
		if q == "quit" {
			i = 100

		}
	}
}
