//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var c byte
	var n, s int
	for{
		fmt.Scanf("%c", &c)
		if c == '2'{
			break
		}else if c == '0' && s == 0{
			n++
			s++
		}else if c == '1'{
			s = 0
		}
	}
	fmt.Println("Sequenze di zeri:", n)
}