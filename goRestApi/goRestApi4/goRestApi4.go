package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type Users struct {

	//It is important to make atributes with capitol first letter!!!
	//`json:` is important if we're sending and reciving json format code
	gorm.Model
	//User_ID  string  `gorm:"primary_key:true"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Name     string  `json:"name"`
	Lastname string  `json:"lastname"`
	Email    string  `json:"email"`
	Roles    []Roles `gorm:"many2many:user_role" json:"roles"`
	//Role     []Roles `gorm:"many2many:Users_Roles;association_foreignkey:RolesID;foreignkey:UserID"`
}
type Roles struct {
	gorm.Model

	Name string `json:"roles"`
}

func InitialMigration() {
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=goLangTest password=password")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database!")
	}

	defer db.Close()

	db.AutoMigrate(&Users{}, &Roles{})

}

func AllUsers(w http.ResponseWriter, r *http.Request) {

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=goLangTest password=password sslmode=disable")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database!")
	}
	defer db.Close()

	var arrayOfUsers []Users
	db.Find(&arrayOfUsers)

	json.NewEncoder(w).Encode(arrayOfUsers)

	fmt.Println("Endpoint Hit: AllUsers")
}

func Createuser(w http.ResponseWriter, r *http.Request) {

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=goLangTest password=password")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database!")
	}
	defer db.Close()

	var newUser Users
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
	}

	json.Unmarshal(reqBody, &newUser)

	newUser.Roles = []Roles{Roles{Name: "USER"}}

	db.Create(&newUser)

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

	var users Users

	db.Where("name=?", name).Find(&users)
	db.Delete(&users)

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
	var users Users
	db.Where("name=?", name).Find(&users)

	users.Email = email

	db.Save(&users)

	fmt.Println("Endpoint Hit: updateUser")
}
