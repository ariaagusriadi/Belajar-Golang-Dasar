package main

import "fmt"

func getAge(age int) int {
	return 2023 - age
}

func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello" + name
	}
}

func main() {
	age := getAge(2002)
	fmt.Println(age)

	result := getHello("Aria")
	fmt.Println(result)
	fmt.Println(getHello(""))
}
