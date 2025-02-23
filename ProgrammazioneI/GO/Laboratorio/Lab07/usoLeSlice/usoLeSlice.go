//Autore: Francesco Corrado

package main

import(
    "fmt"
	"bufio"
	"os"
	"sort"
)

func main(){
    var mySlice[]string
	var newSlice2[]string
	var sli1[]string
	var sli2[] string
	var s3[] string
	var s string
	sli1 = os.Args[1:]
	fmt.Println("\nScrivi almeno 3 parole (seguite da ctrl D)")
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Split(bufio.ScanWords)
	for myScanner.Scan() {
		word := myScanner.Text()
		sli2=append(sli2, word)
	}
	mySlice=append(sli1, sli2...)
	fmt.Println("\n 1.", mySlice)
	sort.Strings(mySlice)
	fmt.Println(" 2.", mySlice)
	mySlice = mySlice[:len(mySlice)-1]
	fmt.Println(" 3.", mySlice)
	mySlice = append(mySlice[:2], mySlice[4:]...)
	fmt.Println(" 4.", mySlice)
	newSlice:=[]string{"aa", "bb", "cc"}
	fmt.Println(" 5.", newSlice)
	mySlice = append(mySlice[:1], append(newSlice, mySlice[1:]...)...)
	fmt.Println(" 6.", mySlice)
	fmt.Print("Scrivi una parola: ")
	fmt.Scan(&s) //Prova ad ottimizzare
	s3 = append(s3, s)
	mySlice = append(mySlice[:2], append(s3, mySlice[2:]...)...)
	fmt.Println(" 7.", mySlice)
	fmt.Print("Scrivi una parola: ")
	fmt.Scan(&s)
	mySlice = append(mySlice, s)
	fmt.Println(" 8.", mySlice)
	s3 = make([]string, 2)
	mySlice = append(mySlice, s3...)
	fmt.Println(" 9.", mySlice)
	s3 = make([]string, 3)
	mySlice = append(mySlice[:3], append(s3, mySlice[3:]...)...)
	fmt.Println("10.", mySlice)
	newSlice2 = append(newSlice2, mySlice...)
	fmt.Println("11.", newSlice2)
	newSlice2 = newSlice2[:len(newSlice2)-1]
	fmt.Println("12.", mySlice, "\n   ", newSlice2)
}