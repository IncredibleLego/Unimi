//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var s string
	s="Garibaldi fu ferito"
	for i:=0; i<len(s); i++{
		c:=rune(s[i])
		if c=='a' || c=='e' || c=='i' || c=='o' {
			c='u'
		}
		fmt.Print(string(c))
	}
	fmt.Println()
}