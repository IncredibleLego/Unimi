//Autore: Francesco Corrado
package main

import "fmt"

func main(){
	var centigradi, fahrenheit float64
	fmt.Println("Programma che converte in gradi Fahrenheit una data temperatura in Celsius")
	ratio:=5./9
	fmt.Print("Inserire la temperatura da convertire: ")
	fmt.Scan(&centigradi)
	fahrenheit=centigradi/ratio+32
	fmt.Println(centigradi, "C =", fahrenheit, "F")
}