//Autore: Francesco Corrado

package main

import(
	"fmt"
	"math"
)

func main(){
	var n int
	fmt.Scan(&n)
	var a[200]int
	for i:=0; i<n; i++{ //Inserisco n valori in a
		fmt.Scan(&a[i])
	}
	s:=0
	for i:=0; i<n; i++{
		s+=a[i]
	}
	media:=float64(s)/float64(n) //Sommo tutti i valori in s e calcolo la media
	sq:=0.0
	for i:=0; i<n; i++{
		sq += (float64(a[i])-media)*(float64(a[i])-media) //Calcolo quanto ogni valore dista dalla media
	}
	sqm:= math.Sqrt(sq/float64(n)) //Più sqm è alto più i valori sono dispersi
	fmt.Println("SQM:", sqm)
}