package main

import "fmt"

func getComplateName() (firstName string, lastName string, age int) {
	firstName = "Aria"
	lastName = "Riadi"
	age = 21
	return
}

func main() {
	a, b, c := getComplateName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
