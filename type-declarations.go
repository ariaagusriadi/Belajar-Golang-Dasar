package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTPAria NoKTP = "12223311122333"
	var marriedStatus Married = false

	fmt.Println(noKTPAria)
	fmt.Println(marriedStatus)
}
