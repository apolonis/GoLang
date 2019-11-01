package main

import "fmt"

func main() {

	go capture()
	go analyze()
	fmt.Scanln()

}
