package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Aria"
	names[1] = "Agus"
	names[2] = "Riadi"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90, 94, 24,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
	fmt.Println(len(names))
}
