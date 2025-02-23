package main
import "fmt"
/*Problema: Scrivere un programma Go no_out.go che dichiara una variabile int e una costante = 10,
assegna alla variabile un valore doppio della costante, poi somma 1 alla variabile, ma non stampa niente. */

func main() {
	var a int
	const b = 10
	a = b*2
	a = a + 1
	fmt.Println("A =", a, " B =", b)
}