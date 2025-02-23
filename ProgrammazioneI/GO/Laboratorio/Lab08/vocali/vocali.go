//Autore: Francesco Corrado

package main

import(
	"fmt"
	"bufio"
	"os"
	"sort"
)

func contaVocali(s string, vocali map[rune]int) map[rune]int {
	for i:=0; i<len(s); i++{
		c := rune(s[i])
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'{
			vocali[c] += 1
		}
	}
	return vocali
}

func main(){
	var m map[rune]int
	m = make(map[rune]int)
	_=m
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
    	word := scanner.Text()
		m = contaVocali(word, m)
  	}
	keys := make([]string, 0, len(m))
	for k := range m{
		f := string(k)
        keys = append(keys, f)
    }
    sort.Strings(keys)
    for _, k := range keys {
		g := []rune(k)
        fmt.Println("\t", k, ":", m[g[0]])
    }
}