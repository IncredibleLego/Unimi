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
	var out[]string
	var i int
	myScanner := bufio.NewScanner(os.Stdin)
	for myScanner.Scan() {
		line := myScanner.Text()
		if line == "stop"{
			break
		}
		sli=append(sli, line)
	}
	for y:=0; y<len(sli); y++{
		x := string(sli[y])
		i=0
		out = []string{}
		for c:=0; c<len(x); c++{
			if unicode.IsPunct(rune(x[i])){
				c++
			}else{
				out=append(out, string(x[i]))
			}
			i++
		}
		for r:=0; r<len(out); r++{
			fmt.Print(out[r])
		}
		fmt.Println()
	}
}