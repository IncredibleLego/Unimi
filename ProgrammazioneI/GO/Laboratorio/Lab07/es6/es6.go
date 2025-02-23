//Autore: Francesco Corrado

package main

import(
	"fmt"
	"os"
	"bufio"
	"unicode"
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
		var s string
		s = sli[i]
		for x:=0; x<len(s); x++{
			if unicode.IsDigit(rune(s[x])){
				fmt.Println(sli[i])
				break
			}
		}
	}
}