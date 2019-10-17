package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name  string
	Email string
}

func InitialMigration() {
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=goLangTest password=password")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database!")
	}
	defer db.Close()
	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=goLangTest password=password")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database!")
	}
	defer db.Close()

	var users []User
	db.Find(&users)

	json.NewEncoder(w).Encode(users)

	fmt.Println("Endpoint Hit: AllUsers")
}
func Createuser(w http.ResponseWriter, r *http.Request) {

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=goLangTest password=password")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database!")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]
	db.Create(&User{Name: name, Email: email})

	fmt.Fprintf(w, "New user has been created!")

	fmt.Println("Endpoint Hit: New user has been created")
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User has been deleted!")
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=goLangTest password=password")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database!")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User

	db.Where("name=?", name).Find(&user)
	db.Delete(&user)

	fmt.Println("User has been deleted!")
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "updateUser endpoint hit!")

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=goLangTest password=password")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database!")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]
	var user User
	db.Where("name=?", name).Find(&user)

	user.Email = email

	db.Save(&user)

	fmt.Println("Endpoint Hit: updateUser")
}
