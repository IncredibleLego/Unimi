//Autore: Francesco Corrado

package main

import(
    "fmt"
	"os"
)

func anagramma(s string) []string{ //Provare a fare senza inserire questa funzione qua ma sistemandolo
	if s=="" {
		return []string{""}
	}else{
		var ris []string
		for i:=0; i<len(s); i++{
			primo := rune(s[i])
			resto := s[:i]+s[i+1:]
			anagResto:=anagramma(resto)
			for _, x:= range anagResto{
				ris = append(ris, string(primo)+x)
			}
		}
		return ris
	}
}

func isAnagram(s1, s2 string) bool{
	var anag []string
	anag = anagramma(s2)
	for x:=0; x<len(anag); x++{
		if s1 == anag[x]{
			return true
		}
	}
	return false
}

func main(){
    var p1, p2 string
	p1 = os.Args[1]
	p2 = os.Args[2]
	if isAnagram(p1, p2){
		fmt.Println("p1 e p2 sono anagrammi")
	}else{
		fmt.Println("p1 e p2 non sono anagrammi")
	}
}