package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll(10, 11, 22, 33, 44)
	fmt.Println(total)

	slice := []int{10, 11, 22, 33, 44}
	ttl := sumAll(slice...)
	fmt.Println(ttl)
}
