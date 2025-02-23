//Autore: Francesco Corrado

package main

import(
	"fmt"
	"os"
	"bufio"
)


func main(){
	var sli[]string
	var s int
	myScanner := bufio.NewScanner(os.Stdin)
	for myScanner.Scan() {
		line := myScanner.Text()
		sli=append(sli, line)
	}
	for i:=0; i<len(sli); i++{
		s += len(sli[i])
	}
	media:= s/len(sli)
	for i:=0; i<len(sli); i++{
		if len(sli[i])>media{
			fmt.Print(sli[i], " ")
		}
	}
	fmt.Println()
}