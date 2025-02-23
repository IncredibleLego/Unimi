//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func printBinary(x int){
	if x<2{
		fmt.Print(x)
	}else{
		resto:=x%2
		printBinary(x/2)
		fmt.Print(resto)
	}
}

func main(){
	var n int
	fmt.Print("numero: ")
	fmt.Scan(&n)
	printBinary(n)
	fmt.Println()
}