//Autore: Francesco Corrado

package main

import(
	"fmt"
	"os"
	"bufio"
)


func main(){
	var sli[]string
//	sli=make([]string,0,10)
//	fmt.Println(cap(sli), len(sli))
	myScanner := bufio.NewScanner(os.Stdin)
	for myScanner.Scan() {
		line := myScanner.Text()
		if line == "stop"{
			break
		}
		sli=append(sli, line)
	}
//	fmt.Println(cap(sli), len(sli))
	for i:= len(sli)-1; i >= 0; i--{
		fmt.Println(sli[i])
	}
}