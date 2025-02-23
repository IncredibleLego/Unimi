//Autore: Francesco Corrado

package main

import(
	"fmt"
	"os"
	"bufio"
	"sort"
)


func main(){
	var sli[]string
	myScanner := bufio.NewScanner(os.Stdin)
	for myScanner.Scan() {
		line := myScanner.Text()
		if line == "stop"{
			break
		}
		sli=append(sli, line)
	}
	ord := sli
	sort.Strings(ord)
	for i:=0; i<len(ord); i++{
		fmt.Print(ord[i], " ")
	}
	fmt.Println()
}