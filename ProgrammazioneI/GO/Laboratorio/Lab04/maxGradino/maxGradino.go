//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var c, p, n1, n2, max int
	n1 = 1
	n2 = 1
	p = -1
	max = -1
	for{
		fmt.Scan(&c)
		if c < p{
			break
		}else if c == p{
			n2++
		}else{
			if n1 + n2 > max{
				max = n1 + n2
			}
			n1 = n2
			n2 = 1
		}
		p = c
	}
	fmt.Println("Gradino Max:", max)
}