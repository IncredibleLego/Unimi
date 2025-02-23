//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var g, m int
	fmt.Print("Inserire giorno e mese dell'anno: ")
	fmt.Scan(&g, &m)
	if g>0 && g<32 && m>0 && m<13{
		fmt.Println("data valida")
	}else{
		fmt.Println("Data non valida")
	}
}