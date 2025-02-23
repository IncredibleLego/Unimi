//Autore: Francesco Corrado

package main

import(
	"fmt"
	"os"
)

func main(){
	var s string
	fmt.Print("Inserire il valore")
	fmt.Scan(&s)
	f, err := os.Create(s)
	defer f.Close()
	var b byte[]
	
}