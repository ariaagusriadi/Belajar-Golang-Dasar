package main

import "fmt"

func main() {
	name := "agus riadi"

	if name == "aria" {
		fmt.Println("Hello Aria")
	} else if name == "agus" {
		fmt.Println("Hello Agus")
	} else {
		fmt.Println("Hi, world!")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
