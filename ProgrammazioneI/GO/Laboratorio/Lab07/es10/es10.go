//Autore: Francesco Corrado

package main

import(
	"fmt"
	"os"
	"bufio"
)


func main(){
	var sli[]string
	var max int
	myScanner := bufio.NewScanner(os.Stdin)
	for myScanner.Scan() {
		line := myScanner.Text()
		sli=append(sli, line)
	}
	for i:=0; i<len(sli); i++{
		if len(sli[i])>max{
			max = len(sli[i])
		}
	}
	for i:=0; i<len(sli); i++{
		if len(sli[i]) == max{
			fmt.Print(sli[i], " ")
		}
	}
	fmt.Println()
}