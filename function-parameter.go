package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hi", firstName, lastName)
}

func main() {
	sayHelloTo("Aria", "agus")
	sayHelloTo("Fek", "liao")
}
