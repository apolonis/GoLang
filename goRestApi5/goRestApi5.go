package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name     string
	Lastname string
}
type allUsers []User

var Users = allUsers{
	{

		Name:     "admin",
		Lastname: "admin",
	},
}

//Req mapping home page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func Createuser(w http.ResponseWriter, r *http.Request) {

	var newUser User
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
	}

	json.Unmarshal(reqBody, &newUser)
	json.NewEncoder(w).Encode(newUser)
	fmt.Fprintf(w, "New user has been created!")

	fmt.Println("Endpoint Hit: New user has been created")
}

//Get mapping/get all users
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Users)
}
func handleRequests() {

	router := mux.NewRouter() //router

	router.HandleFunc("/", homePage)                              //req map - home page
	router.HandleFunc("/user/{user}", createUser).Methods("POST") //post/create
	router.HandleFunc("/user", getAllUsers).Methods("GET")        //get all users

	log.Fatal(http.ListenAndServe(":8080", router)) // port
}

func main() {
	handleRequests()

	////// The syntax for creating an endpoint looks like this:
	////// router.HandleFunc("/<your-url>", <function-name>).Methods("<method>")

	// func <your-function-name>(w http.ResponseWriter, r *http.Request)

}
