package main

import(
	"fmt"
	"os"
	"bufio"
	"sort"
	"strings"
)

func stampa(m map[string][]int){
	var  val []string
	for i := range m{
		val = append(val, i)
	}
	sort.Strings(val)
	for _, i := range val{
		if i != ""{
			fmt.Printf("%s: %v", i, m[i])
			fmt.Println()
		}
	}
}


func main() {
	var righe map[string][]int
	righe = make(map[string][]int)
	var indice []int
	n := 1
	fmt.Println("+ (memorizza)\n? (indica n. di riga)\np (stampa tutto)")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		riga := scanner.Text()
		if string(riga[0]) == "+"{
			substring := strings.Split(riga[2:], " ")
			for i:=0; i < len(substring); i++{
				indice = righe[substring[i]]
				indice = append(indice, n)
				righe[string(substring[i])] = indice
			}
			n++
		}else if string(riga[0]) == "?"{
			substring := riga[2:]
			fmt.Println("stringa:", substring)
			fmt.Println("righe", righe[substring])
		}else{
			stampa(righe)
		}
	}
}