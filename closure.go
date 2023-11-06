package main

import "fmt"

func main() {
	name := "aria"
	counter := 0

	increment := func() {
		name := "agus"
		fmt.Println("increment")
		fmt.Println(name)
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
