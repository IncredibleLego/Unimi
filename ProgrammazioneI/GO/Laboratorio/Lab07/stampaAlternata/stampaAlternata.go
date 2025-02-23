//Autore: Francesco Corrado

package main

import(
	"fmt"
	"os"
	"bufio"
)

func main(){
	var righe []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		riga := scanner.Text()
		righe = append(righe, riga)
	}
	for i:=0; i < len(righe); i++{
		if i % 2 != 0{
			fmt.Println(righe[i])
		}
	}
	for i:=0; i < len(righe); i++{
		if i % 2 == 0{
			fmt.Println(righe[i])
		}
	}
}