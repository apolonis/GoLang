package main

import "fmt"

//USER WITH ROLES AS ATRIBUTES, WITH ARRAYS AND HOW TO GET USER BY CUSTOM ATRIBUTES IN GO LANG

type role struct {
	imeRole string
}

type user struct {
	name     string
	lastname string
	//roles of type role
	roleUsera role
}

func main() {

	//here we're creating array and telling them how manny elements is going to have
	var arrayOfUsers [3]user

	r1 := role{"Admin"}
	r2 := role{"User"}

	u1 := user{"User", "Useric", r1} // u:= user{name:"User", lastname:"Useric", r1}
	u2 := user{"User2", "Useric", r2}
	u3 := user{"User3", "Useric", r2}

	fmt.Println(u1)

	fmt.Println(u2)
	fmt.Println(u3)
	fmt.Println("-----------------")
	u1.name = "User1New" //This is like setter

	arrayOfUsers[0] = u1
	arrayOfUsers[1] = u2
	arrayOfUsers[2] = u3

	fmt.Println(u1)

	fmt.Println(u2)
	fmt.Println(u3)
	fmt.Println("-----------------")
	for i := range arrayOfUsers { // - range arrayOfUsers is like in java -arrayOfUsers.size()

		if arrayOfUsers[i].lastname == "Useric" { //Like getter
			fmt.Println(arrayOfUsers[i])
		}
		i++
	}

	fmt.Println(u1.roleUsera) //This is like getter 'u1.roleUsera'
	fmt.Println("-----------------")
	fmt.Println("Nice job :)")

}
