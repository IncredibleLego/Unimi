package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"sort"
)

func main(){
	sinonimi := make(map[string][]string)
	if len(os.Args) < 3{
		fmt.Println("2 file names, please")
		os.Exit(0)
	}
	nFile1 := os.Args[1]
	nFile2 := os.Args[2]
	file1, err1 := os.Open(nFile1)
	if err1 != nil{
		fmt.Println("file 1 not found")
		os.Exit(0)
	}
	file2, err2 := os.Open(nFile2)
	if err2 != nil{
		fmt.Println("file 2 not found")
		os.Exit(0)
	}
	scanner1 := bufio.NewScanner(file1)
	for scanner1.Scan(){
		riga := scanner1.Text()
		parole := strings.Split(riga, " ")
		sin := parole[1:]
		indice := parole[0]
		indice = indice[:len(indice)-1]
		sinonimi [indice] = sin
	}
	scanner2 := bufio.NewScanner(file2)
	scanner2.Split(bufio.ScanWords)
	for scanner2.Scan(){
		parola := scanner2.Text()
		sin, ok := sinonimi[parola]
		if ok{
			var o string
			for i:=0; i < len(sin)-1; i++{
				s := sin[i]
				for j:=0; j < len(s)-1; j++{
					o += string(s[j])
				}
				sin[i] = o
				o = ""
			}
			sort.Strings(sin)
			out := fmt.Sprintf("%s%s", parola, sin)
			fmt.Print(out, " ")
		}else{
			fmt.Print(parola, " ")
		}
	}
	fmt.Println()
	file1.Close()
	file2.Close()
}