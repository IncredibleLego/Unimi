//Autore: Francesco Corrado

package main

import(
	"fmt"
	"os"
	"strconv"
	"math"
)

func isSquare(n int) bool{
	var num int
	f := math.Sqrt(float64(n))
	num = int(f)
	if f-float64(num) == 0{
		return true
	}
	return false
}

func main(){
	var num []string
	var n int
	num = os.Args[1:]
	for i:=0; i < len(num); i++{
		n, _ = strconv.Atoi(num[i])
		if isSquare(n){
			fmt.Println(n)
		}
	}
}