//Autore: Francesco Corrado

package main

import(
	"fmt"
	"strings"
	"strconv"
)

func main(){
	var s string
	var g, m, a int
	fmt.Print("Inserisci la data: ")
	fmt.Scan(&s)
	g, m, a = parseDate(s)
	c:=daysFromEpoch(g, m, a)
	fmt.Println("Sono passati", c, "giorni")
	
}

func daysFromEpoch (g, m, a int) int{
	var count int
	for x:=1970; x<a; x++{
		count+=365
		if isLeap(x){
			count++
		}
	}
	for x:=1; x<m; x++{
		if x == 11 || x == 4 || x == 6 || x == 9 {
			count+=30
		}else if x == 2{
			if isLeap(a){
				count+=2
			}else{
				count+=28
			}
		}else{
			count+=31
		}
	}
	count+=m
	return count
}

func isLeap (a int) bool{ //Funzione che calcola se l'anno e bisestile
	return (a%4==0 && a%100!=0 || a%400 ==0)
}

func parseDate (s string) (int, int, int){ //Funzione che converte un anno da stringa a piu variabili int
	b1:=strings.Index(s, "/")
	b2:=strings.Index(s[b1+1:], "/")+ b1+1
	sg:= s[:b1]
	sm:= s[b1+1:b2]
	sa:= s[b2+1:]
	g, _ := strconv.Atoi(sg)
	m, _ := strconv.Atoi(sm)
	a, _ := strconv.Atoi(sa)
	return g, m, a
}