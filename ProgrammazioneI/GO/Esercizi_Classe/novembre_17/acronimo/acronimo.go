//Autore: Francesco Corrado

package main

import(
	"fmt"
	"bufio"
	"os"
)

func acronimo (s []string) string{
	var t string
	for _, x:= range s{
		t+=string(x[0])
	}
	return t
}

func main(){
	var x []string
	scanner:=bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan(){
		x = append(x, scanner.Text())
	}
	t:=acronimo(x)
	fmt.Println(t)
}

/*
func main(){
	var x []string
	x = []string{"Giovanni", "cade", "dalla", "finestra"}
	t:=acronimo(x)
	fmt.Println(t)
}
*/