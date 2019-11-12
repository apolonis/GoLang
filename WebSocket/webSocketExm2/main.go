package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "index.html")
}

func countriesAll(w http.ResponseWriter, r *http.Request) {
	ws, err := Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "DEBUG: in function coutries all %+v\n", err)
	}

	go Writer(ws)
}

func setUpRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/countriesAll", countriesAll)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Programm is running")
	setUpRoutes()

	// country, err := getCountries()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for i := 0; i < len(country); i++ {
	// 	fmt.Println(country[i])
	// }
}
