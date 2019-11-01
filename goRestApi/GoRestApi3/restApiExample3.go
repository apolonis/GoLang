package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//BUILDING REST API APP FOR SAVING USERS IN ARRAY
type role struct {
	Name string
}
type user struct {

	//It is important to make atributes with capitol first letter!!!
	//`json:` is important if we're sending and reciving json format code

	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Role     role
}

type allUsers []user

var users = allUsers{
	{
		Username: "admin",
		Password: "admin",
		Name:     "admin",
		Lastname: "admin",
		Email:    "admin@admin.com",
		Role:     role{"ADMIN"},
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
	newUser.Role = role{"USER"}
	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
	fmt.Println("Endpoint Hit: User has been created")
}

//Delete mapping
func deleteUser(w http.ResponseWriter, r *http.Request) {
	userName := mux.Vars(r)["name"]
	for i, singleUser := range users {
		if singleUser.Username == userName {
			users = append(users[:i], users[i+1:]...)
			fmt.Fprintf(w, "The user with name %v has been deleted successfully", userName)
		}
	}
}

//Update mapping
func updateUser(w http.ResponseWriter, r *http.Request) {
	//"name"--- mux.Vars(r)["name"] - find any atribute, in bellow 'singleUser.atribute' == mux.Vars(r)["name"]
	//Thats like pathVariable in java SpringBoot
	userName := mux.Vars(r)["name"]
	var updatedUser user
	//This is like requestBody in java SpringBoot
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Enter valid informations")
	}

	json.Unmarshal(reqBody, &updatedUser)

	for i, singleUser := range users {

		if singleUser.Username == userName {
			updatedUser.Role = role{"USER"}
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

	////// The syntax for creating an endpoint looks like this:
	////// router.HandleFunc("/<your-url>", <function-name>).Methods("<method>")

	// func <your-function-name>(w http.ResponseWriter, r *http.Request)

}
