//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func freq(s string) map[rune]int {
	var res map[rune]int
	res = make(map[rune]int)
	for _, r:= range s{
		res[r]++
	}
	return res
}

func main(){
	var str string
	fmt.Print("Inserire una stringa: ")
	fmt.Scan(&str)
	fmt.Println(freq(str))
}