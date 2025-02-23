package main

import "fmt"

func main() {
	var dividendo, divisore int
	fmt.Println("fornire due numeri (int) di cui si vuole calcolare quoziente e resto")
	fmt.Print("dividendo: ")
	fmt.Scan(&dividendo)
	fmt.Print("divisore: ")
	fmt.Scan(&divisore)
	quoziente:=dividendo/divisore
	resto:=dividendo%divisore
	fmt.Println("quoziente =", quoziente)
	fmt.Println("resto =", resto)
}
