//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	for{
		var o, m int
		fmt.Print("ore e minuti:")
		fmt.Scan(&o, &m)
		if o>=0 && o<=23 && m>=0 && m<=59{
			fmt.Println("valori validi")
			break
		}
	}
}