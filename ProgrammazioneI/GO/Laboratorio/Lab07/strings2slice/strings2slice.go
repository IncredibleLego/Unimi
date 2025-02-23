//Autore: Francesco Corrado

package main

import(
	"fmt"
	"strings"
	"os"
)

func main(){
	var s string
	s = os.Args[1]
	sli1 := strings.Split(s, ",")
	fmt.Println("Sli1: ", sli1)
	sli2 := strings.SplitAfter(s, ",")
	fmt.Println("Sli2: ", sli2)
	s = os.Args[2]
	sli3 := strings.Fields(s)
	fmt.Println("Sli3: ", sli3)
	s = os.Args[3]
	fmt.Print("Codice ascii numero ", s, ": ")
	for i:=0; i < len(s); i++{
		fmt.Print(s[i], " ")
	}
	fmt.Println()
	s = os.Args[4]
	fmt.Print("Codice unicode delle lettere '", s, "': ")
	for i:=0; i < len(s); i++{
		fmt.Print(rune(s[i]), " ")
	}
	fmt.Println()
	s = os.Args[5]
	sli4 := strings.Split(s, ":")
	fmt.Println("Sli4: ", sli4)
}