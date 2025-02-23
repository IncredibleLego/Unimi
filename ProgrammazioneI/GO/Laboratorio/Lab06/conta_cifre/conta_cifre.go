//Autore: Francesco Corrado

package main

import(
	"fmt"
	"unicode"
	"bufio"
	"os"
	"strconv"
)

func contaCifre(s string, numCifre *[10]int) (haCifre bool){
	var c int
	for i:=0; i<len(s); i++{
		if unicode.IsDigit(rune(s[i])){
			n,_ := strconv.Atoi(string(s[i]))
			numCifre[n]++
			c = 1
		}
	}
	if c == 0{
		return false
	}else{
		return true
	}
}

func main(){
	var s1[]string
	var occ, num [10]int
	var co int
	for i:=0; i < 10; i++{
		occ[i] = 0
	}
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Split(bufio.ScanWords)
	for myScanner.Scan() {
		word :=myScanner.Text()
		if word == "stop"{
			break
		}
		s1=append(s1, word)
	}
	for i:=0; i<len(s1); i++{
		if contaCifre(string(s1[i]), &occ){
			co++
		}
	}
	fmt.Println(co, "stringhe con cifre")
	for i:=0; i < 10; i++{
		num[i] = i
	}
	fmt.Println(num)
	fmt.Println(occ)
}