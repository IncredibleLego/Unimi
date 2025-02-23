package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)
func cancella(lista []string) []string{
	var output []string
	for i:=0; i < len(lista); i++{
		n, err := strconv.Atoi(lista[i])
		if err == nil{
			i = i + n
		}else{
			output = append(output, lista[i])
		}
	}
	return output
}

func main(){
	var slice []string
	if len(os.Args) != 2 {
		fmt.Println("Fornire 1 nome di file")
		os.Exit(0)
	}
	nFile := os.Args[1]
	f, err := os.Open(nFile)
	if err != nil{
		fmt.Println("File non accessibile")
		os.Exit(0)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan(){
		parola := scanner.Text()
		slice = append(slice, parola)
	}
	fmt.Println(slice)
	sli2 := cancella(slice)
	fmt.Println(sli2)
	f.Close()
}