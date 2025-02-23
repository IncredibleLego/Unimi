//Autore: Francesco Corrado

package main

import(
	"fmt"
	"bufio"
	"os"
	"unicode"
)

func hasUpper (s string) bool{
	for i:=0; i < len(s); i++{
		r := rune(s[i])
		if unicode.IsUpper(r){
			return true
		}
	}
	return false
}

func firstUpper (s string) int{
	for i:=0; i < len(s); i++{
		r := rune(s[i])
		if unicode.IsUpper(r){
			return i
		}
	}
	return -1
}

func lastUpper (s string) int{
	for i:=len(s)-1; i >= 0; i--{
		r := rune(s[i])
		if unicode.IsUpper(r){
			return i
		}
	}
	return -1
}

func countDigitsLettersOthers (s string) (int, int, int){
	var c, l, a int
	for i:=0; i < len(s); i++{
		r := rune(s[i])
		if unicode.IsNumber(r){
			c++
		}else if unicode.IsLetter(r){
			l++
		}else{
			if !unicode.IsSpace(r){
				a++
			}
		}
	}
	return c, l, a

}

func isPalindrome (s string) bool{
	var stringaReverse string
	for i:=len(s)-1; i >= 0; i--{
		stringaReverse += string(s[i])
	}
	if stringaReverse == s{
		return true
	}else{
		return false
	}
}


func main(){
	var a, b, c int
	scanner:= bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		S := scanner.Text()
		if hasUpper(S){
			fmt.Println("ha maiuscole")
		}else{
			fmt.Println("non ha maiuscole")
		}
		fmt.Println("prima maiuscola in posizione", firstUpper(S))
		fmt.Println("ultima maiuscola in posizione", lastUpper(S))
		a, b, c = countDigitsLettersOthers(S)
		fmt.Println("cifre, lettere, altro:", a, b, c)
		if isPalindrome(S){
			fmt.Println("palindroma")
		}else{
			fmt.Println("non palindroma")
		}
	}
}