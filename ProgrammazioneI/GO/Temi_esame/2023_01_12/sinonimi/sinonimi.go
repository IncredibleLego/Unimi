package main

import(
    "fmt"
	"os"
	"bufio"
	"strings"
	"sort"
	"unicode"
)
/*
Funzione che restituisce la stringa da mettere dopo la parola, se la parola ha un sinonimo
trova in sin i vari valori da inserire li trasforma in stringa e li salva, altrimenti
restituisce uno spazio per le parole normali
*/

func trovaSinonimo (par string, sin []string) string{
	for i:=0; i < len(sin); i++{
		var s, sinonim, out, output string
		var ind int
		s = sin[i]
		ind = strings.Index(s, ":")
		sinonim = s[:ind]
		if par == sinonim{
			o := strings.Split(s[ind+1:], ",")
			sort.Strings(o)
			out = "["
			for j:=0; j < len(o); j++{
				out += o[j]
			}
			out += "] "
			for j, r := range out{
				if unicode.IsSpace(r){
					output += out[j+1:]
					break
				}else{
					output += string(out[j])
				}
			}
			return output
		}
	}
	return " "
}

func main(){
	var sinonimi[]string
	var str string
	if len(os.Args) < 3{
		fmt.Println("2 file names, please")
	}else{
		nf1 := os.Args[1]
		nf2 := os.Args[2]
	
		f1, err1 := os.Open(nf1)
		if err1 == nil{
			scanner1 := bufio.NewScanner(f1)
			for scanner1.Scan(){
				riga := scanner1.Text()
				sinonimi = append(sinonimi, riga)
			}
		}else{
			fmt.Println("file 1 not found")
		}
		f2, err2 := os.Open(nf2)
		if err2 == nil{
			scanner2 := bufio.NewScanner(f2)
			scanner2.Split(bufio.ScanWords)
			for scanner2.Scan(){
				parola := scanner2.Text()
				str = trovaSinonimo(parola, sinonimi)
		
				fmt.Print(parola, str)
			}
			fmt.Println()
		}else{
			fmt.Println("file 2 not found")
		}
		f1.Close()
		f2.Close()
	}
}