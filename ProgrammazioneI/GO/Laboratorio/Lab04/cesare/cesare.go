//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var k int
	var s, f string
	var c byte
	fmt.Scan(&k)
	fmt.Scan(&s)
	for i:=0; i<len(s); i++{
		c = byte(s[i]) + byte(k)
		if c > 122{
			for{
				c = c - 26
				if c >= 97 && c <=122{
					break
				}
			}
		}
		f = f + string(c)
	}
	fmt.Println(f)
}