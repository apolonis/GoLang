package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

// main
func main() {
	// creates a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()

	// out of the box fsnotify can watch a single file, or a single directory
	err = watcher.Add("/home/marko/ewbMarko/GoLang/GoExm Functionality/worker0/exm.txt")

	if err != nil {
		fmt.Println("ERROR", err)
	}

	//
	done := make(chan bool)

	//
	go func(done chan bool) {
		for {
			select {
			// watch for events
			case event, ok := <-watcher.Events:
				fmt.Println("EVENT:", event, ok)

			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}(done)

	<-done
}
