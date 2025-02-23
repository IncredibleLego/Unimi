//Autore: Francesco Corrado

package main

import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	var s, o string
	var c, l int
	_ = s
	scanner:= bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan(){
		s:= scanner.Text()
		if c == 0{
			o = s
			c++
		}else{
			if len(o) > len(s){
				l = len(s)
			}else{
				l = len(o)
			}
			for i:=0; i<l; i++{
				if s[i] < o[i]{
					o = s
				}
			}
		}
	}
	fmt.Println("\n", o)
}

//Soluzione prof
func main(){
	var minimo string
	primaVolta := true
	scanner:= bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan(){
		word:=scanner.Text()
		if primaVolta || word < minimo {
			minimo = word
			primaVolta = false
		}
	}
}