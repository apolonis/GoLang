package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// type role struct {
// 	imeRole string
// }

type user struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Role     string `json:"role"`
	//roles of type role
	//roleUsera role
}
type allUsers []user

var users = allUsers{
	{
		Name:     "User0",
		Lastname: "Useric0",
		Role:     "User0",
	},
}

//Req mapping home page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

//Post mapping/ create new user
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser user
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
	}

	json.Unmarshal(reqBody, &newUser)
	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
	fmt.Println("Endpoint Hit: User has been created")
}

//Delete mapping
func deleteUser(w http.ResponseWriter, r *http.Request) {
	userName := mux.Vars(r)["name"]
	for i, singleUser := range users {
		if singleUser.Name == userName {
			users = append(users[:i], users[i+1:]...)
			fmt.Fprintf(w, "The user with name %v has been deleted successfully", userName)
		}
	}
}

//Update mapping
func updateUser(w http.ResponseWriter, r *http.Request) {
	userName := mux.Vars(r)["name"]
	var updatedUser user

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter valid informations")
	}
	json.Unmarshal(reqBody, &updatedUser)

	for i, singleUser := range users {
		if singleUser.Name == userName {
			singleUser = updatedUser
			users = append(users[:i], singleUser)
			json.NewEncoder(w).Encode(singleUser)
		}
	}
}

//Get mapping/get all users
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}
func handleRequests() {

	router := mux.NewRouter() //router

	router.HandleFunc("/", homePage)                                       //req map - home page
	router.HandleFunc("/user/create", createUser).Methods("POST")          //post/create
	router.HandleFunc("/user", getAllUsers).Methods("GET")                 //get all users
	router.HandleFunc("/user/update/{name}", updateUser).Methods("PUT")    // update mapping
	router.HandleFunc("/user/delete/{name}", deleteUser).Methods("DELETE") // delete mapping

	log.Fatal(http.ListenAndServe(":8080", router)) // port
}

func main() {
	handleRequests()

}
