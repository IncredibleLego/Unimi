//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var c1, c2, c3 int
	fmt.Println("Programma che rileva i guasti a dei componenti")
	fmt.Print("Componente 1, guasto rilevato? (0 per ok) ")
	fmt.Scan(&c1)
	fmt.Print("Componente 2, guasto rilevato? (0 per ok) ")
	fmt.Scan(&c2)
	fmt.Print("Componente 2, guasto rilevato? (0 per ok) ")
	fmt.Scan(&c3)
	if c1 == 0 && c2 == 0 && c3 == 0 {
		fmt.Println("nessun guasto rilevato")
	}else{
		fmt.Println("guasti rilevati a:")
		if c1 != 0{
			fmt.Println("caricamento acqua")
		}
		if c2 != 0{
			fmt.Println("scarico acqua")
		}
		if c3 != 0{
			fmt.Println("sistema di riscaldamento")
		}
	}
}