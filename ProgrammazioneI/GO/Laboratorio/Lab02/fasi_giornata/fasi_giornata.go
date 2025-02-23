//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main() {
	var o int
	fmt.Println("Programma che data un ora stabilisce la fascia oraria in cui si trova")
	fmt.Print("Inserire l'ora: ")
	fmt.Scan(&o)
	if o>=0 && o<=6 {
		fmt.Println("notte")
	}else if o>=7 && o<=13{
		fmt.Println("mattino")
	}else if o>=14 && o<=18{
		fmt.Println("pomeriggio")
	}else if o>=19 && o<=23{
		fmt.Println("sera")
	}else{
		fmt.Println("orario non valido")
	}
}