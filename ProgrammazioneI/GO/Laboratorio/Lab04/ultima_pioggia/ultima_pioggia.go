//Autore: Francesco Corrado

package main

import(
	"fmt"
	"io"
)

func main(){
	var n, c, last int
	c = 1
	for{
	//	fmt.Scan(&n)
		_, err := fmt.Scan(&n)
			if err == io.EOF {
				break
			}
		if n > 0{
			last = c
		}
		c++
	/*	if n == 9{
			break
		} */
	}
	fmt.Println("L'ultimo giorno che ha piovuto è stato il giorno ", last)
}