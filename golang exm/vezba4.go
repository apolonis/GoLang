package main

import "fmt"

//THIS EXAMPLE IS FOR MAP's, HOW TO USE THEM IN 2 WAYS

type car struct {
	name string
	tip  string
}

// func main(){
// 	v:=car{"Ferrari", "Enzo"}

// 	var mapa = map[string]car{

// 	"Car":{v.name,v.tip},

// 	}

// 	fmt.Println(mapa)

// }

func main() {

	v := car{"Ferrari", "Enzo"}

	mapa := make(map[string]car)

	mapa["Car"] = v

	fmt.Println(mapa)

}
