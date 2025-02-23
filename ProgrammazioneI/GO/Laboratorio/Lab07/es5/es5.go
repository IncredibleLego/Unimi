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
		var s string
		s = sli[i]
		if string(s[0]) == "a" || string(s[0]) == "e" || string(s[0]) == "i" || string(s[0]) == "o" || string(s[0]) == "u" || string(s[0]) == "A" || string(s[0]) == "E" || string(s[0]) == "I" || string(s[0]) == "O" || string(s[0]) == "U"{
			fmt.Println(sli[i])
		}
	}
}