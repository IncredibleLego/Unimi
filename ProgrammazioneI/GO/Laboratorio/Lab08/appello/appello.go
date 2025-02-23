//Autore: Francesco Corrado

package main

import(
	"fmt"
	"os"
	"slices"
)

func main(){
	if len(os.Args) < 2{
		fmt.Println("not enough data")
		os.Exit(0)
	}
	nomi := os.Args[1:]
	slices.Sort(nomi)
	fmt.Println(nomi)
}