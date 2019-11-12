package main

import (
	"fmt"
)

//sorting letters
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(Reverse("ukuk elel"))
	fmt.Println(Reverse("): lla meht fo tseb eht si avaj tub gnitseretni si egaugnal oG"))
}
