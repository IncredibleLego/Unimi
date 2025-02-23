//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var x, y float64
	fmt.Scan(&x, &y)
	if x>=0 {
		if y>=0{
			fmt.Println("I quadrante")
		}else{
			fmt.Println("IV quadrante")
		}
	}else{
		if y>=0{
			fmt.Println("II quadrante")
		}else{
			fmt.Println("III quadrante")
		}
	}
}