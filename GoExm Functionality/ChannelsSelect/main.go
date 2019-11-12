package main

import (
	"fmt"
)

//Example for select channels

//select is working for sending as well as reciving

// func cha(In <-chan int, a, b chan int) {
// 	for data := range In { //Recive until closed
// 		select { //Send to first non - blocking channel
// 		case a <- data:
// 		case b <- data:
// 		}
// 	}
// }

// func cha2(inputA, inputB <-chan int, outputA, outputB chan int) {
/////////variable declaration left out for readability, need to declare
// 	for {

// 		select { //recive from first non-blocking
// 		case data, more = <-inputA:
// 		case data, more = <-inputB:
// 		}

// 		if !more {
// 			return
// 		}

// 		select { // send to first non-blocking
// 		case outputA <- data:
// 		case outputB <- data:
// 		}

// 	}

// }

// func quitChannel(Quit <-chan int, inpA, inpB, outA, outB chan int) {
// 	//variable declaration left out for readability

// 	for {

// 		select {

// 		case data = <-inpA: //remember : close generates a message
// 		case data = <-inpB: //Actualy this is an anti-pattern
// 			//but u can argue that quit acts as a delegate
// 		case <-Quit:
// 			close(inpA)
// 			close(inpB)

// 			Fanout(inpA, outA, outB) //Flush the remaining data
// 			Fanout(inpB, outA, outB)
// 			return

// 		}

// 	}
// }

func main() {
	fmt.Println("Example for select channels")
}
