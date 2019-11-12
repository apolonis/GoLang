package main

import (
	"fmt"
	"sync"
	"time"
)

func f1(ch chan string) {
	return

	for {
		fmt.Println("f1 start")

		fmt.Println("f1 send")
		ch <- "Hello"

		msg := <-ch
		fmt.Println("f1 recv:", msg)

		time.Sleep(2 * time.Second)
		fmt.Println("f1 end")
	}
}

func f2(ch chan string) {
	return

	for {
		fmt.Println("f2 start")

		msg := <-ch
		fmt.Println("f2 recv:", msg)

		fmt.Println("f2 send")
		ch <- "World"

		time.Sleep(2 * time.Second)
		fmt.Println("f2 end")
	}
}

func g1(mutex *sync.Mutex) {
	for {
		mutex.Lock()
		fmt.Println("g1 start")
		time.Sleep(2 * time.Second)
		fmt.Println("g1 end")
		mutex.Unlock()
	}
}

func g2(mutex *sync.Mutex) {
	for {
		mutex.Lock()
		fmt.Println("g2 start")
		time.Sleep(2 * time.Second)
		fmt.Println("g2 end")
		mutex.Unlock()
	}
}

func main() {
	ch := make(chan string)
	go f1(ch)
	go f2(ch)

	mutex := &sync.Mutex{}
	go g1(mutex)
	go g2(mutex)
	fmt.Scanln()
}
