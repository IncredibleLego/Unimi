//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var a, b, MCD int
	for{
		fmt.Scan(&a, &b)
		if b <= a{
			break
		}else{
			fmt.Println("Inserire un valore b <= a")
		}
	}
	for{
		if b == 0{
			MCD = a
			break
		}else{
			r := a%b
			a=b
			b=r
		}
	}
	fmt.Print(MCD, "\n")
}