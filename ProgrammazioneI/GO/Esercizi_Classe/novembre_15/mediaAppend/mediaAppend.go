//Autore: Francesco Corrado

package main

import(
	"fmt"
)


func main(){
	var a []int
	for{
		var x int
		fmt.Scan(&x)
		if x == 0{
			break
		}
		a = append(a, x)
	}
}