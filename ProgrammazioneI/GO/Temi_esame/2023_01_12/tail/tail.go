package main

import(
	"fmt"
	"os"
	"strconv"
	"bufio"
)

func main(){
	var N int
	var s string
	var sli []string
	N, _ = strconv.Atoi(os.Args[1])
	s = os.Args[2]
	f, _ := os.Open(s)
	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		riga:= scanner.Text()
		if len(sli) < N{
			sli = append(sli, riga)
		}else{
			sli = append(sli[1:], riga)
		}
	}
	for i:=0; i<len(sli); i++{
		fmt.Println(sli[i])
	}
	f.Close()
}