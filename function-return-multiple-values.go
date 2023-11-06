package main

import "fmt"

func getFullName() (string, string, string) {
	return "aria", "agus", "riadi"
}

func main() {
	firstName, _, lastName := getFullName()
	fmt.Println(firstName)
	//fmt.Println(midlleName)
	fmt.Println(lastName)
}
