package main

import "fmt"

func main() {
	name := "Eko"
	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Kenalan Donk")
	}

	//switch length := len(name); length > 5 {
	//case true:
	//	fmt.Println("Nama terlalu panjang")
	//case false:
	//	fmt.Println("Nama Sudah Benar")
	//}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan Panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
