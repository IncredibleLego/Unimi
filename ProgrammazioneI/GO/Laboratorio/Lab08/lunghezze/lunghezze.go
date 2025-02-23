//Autore: Francesco Corrado

package main

import(
	"fmt"
	"bufio"
	"os"
	"sort"
)

func contaLunghezza() (map[int]int, int, int){
	var mi, ma int
	var result map[int]int
	result = make(map[int]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	for scanner.Scan() {
		l := len(scanner.Text())
		if l < mi{
			mi = len(scanner.Text())
		}else if l > ma{
			ma = len(scanner.Text())
		}
		result[l]++
	}
	return result, mi, ma
}


func main(){
	m, min, max := contaLunghezza()
	fmt.Println("Min =", min, "Max =", max)
	fmt.Println("METODO DI STAMPA 1: MAPPA INTERA")
	fmt.Println(m)
	fmt.Println("METODO DI STAMPA 2: ELEMENTI IN ORDINE SPARSO")
	for k, v := range m{
		fmt.Printf("%10d\t%5d\n", k, v)
	}
	fmt.Println("METODO DI STAMPA 3: ORDINE CRESCENTE CHIAVI NON PRESENTI INCLUSE")
	keys := make([]int, 0, len(m))
	for k:=min; k <=max; k++{
		keys = append(keys, k)
	}
	sort.Ints(keys) //Se avessi avuto una string in make[]string, 0, len(m) qui uso sort.Strings
    for _, k := range keys {
        fmt.Println("\t", k, ":", m[k])
    }
	fmt.Println("METODO DI STAMPA 4: ORDINE CRESCENTE CHIAVI NON PRESENTI ESCLUSE")
	keys = make([]int, 0, len(m))
	for k := range m{
        keys = append(keys, k)
    }
    sort.Ints(keys)
    for _, k := range keys {
        fmt.Println("\t", k, ":", m[k])
    }
}