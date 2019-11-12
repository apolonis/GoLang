package main

import "fmt"

type Role struct {
	name string
}

type User struct {
	name     string
	lastname string
	role     Role
}

func main() {

	fmt.Println("Welcome!")
	fmt.Println("-----------------")
	r1 := Role{"Admin"}
	r2 := Role{"User"}

	u1 := User{"User1", "Useric1", r1}
	u2 := User{"User2", "Useric2", r2}

	fmt.Println(u1)
	fmt.Println(u2)
	//	fmt.Println("-----------------")
	// mapa := make(map[string]User)
	// mapa2 := make(map[string]User)
	// mapa["User1"] = u1
	// mapa2["User2"] = u2

	// fmt.Println(mapa)
	// fmt.Println(mapa2)

	//	userMap := make(map[string]User)

	userMap := map[string]User{

		"User1": u1,
		"User2": u2,
	}
	fmt.Println("-----------------")
	fmt.Println(userMap)
	fmt.Println("-----------------")

	//This works like empty contstructor
	userr := User{}

	fmt.Println(userr)
	userr.name = "Userr"
	userr.lastname = "Usericc"
	userr.role = Role{"User"}

	fmt.Println(userr)
	fmt.Println("-----------------")
	var arrayOfUsers = []User{u1, u2, userr}

	for i := range arrayOfUsers {
		fmt.Println(arrayOfUsers[i])
	}
	fmt.Println("-----------------")

	for i := range arrayOfUsers {

		if arrayOfUsers[i].role.name == "User" {
			fmt.Println(arrayOfUsers[i])
		}

	}

}
