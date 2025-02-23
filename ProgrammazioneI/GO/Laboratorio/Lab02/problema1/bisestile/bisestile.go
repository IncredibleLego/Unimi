//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var x int
	for{
		fmt.Print("Inserire un anno (positivo): ")
		fmt.Scan(&x)
		if x > -1 {
			break
		}else {
			fmt.Println("L'anno inserito è inferiore a 0, riprovare")
		}
	}
	if ((x%4==0 && x%100!=0) || x%400==0) && x>1581{
		fmt.Println("bisestile")
	}else{
		fmt.Println("non bisestile")
	}
}