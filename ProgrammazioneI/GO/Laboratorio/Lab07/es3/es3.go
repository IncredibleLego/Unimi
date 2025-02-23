//Autore: Francesco Corrado

package main

import(
	"fmt"
	"os"
	"bufio"
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
	for i:=0; i<len(sli); i++{
		if len(sli[i])% 2 == 0{
			fmt.Println(sli[i])
		}
	}
}