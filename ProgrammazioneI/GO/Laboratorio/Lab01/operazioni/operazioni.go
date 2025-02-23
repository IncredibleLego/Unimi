package main

import "fmt"

func main() {
	var x1, x2 float64
	fmt.Print("fornire due numeri (float): ")
	fmt.Scan(&x1, &x2)
	somma:=x1+x2
	differenza:=x1-x2
	prodotto:=x1*x2
	quoziente:=x1/x2
	fmt.Println(x1, "+", x2, "=", somma)
	fmt.Println("differenza =", differenza)
	fmt.Println("il prodotto di", x1, "e", x2, "dà", prodotto)
	fmt.Println(x1, "/", x2, "=", quoziente)
}
