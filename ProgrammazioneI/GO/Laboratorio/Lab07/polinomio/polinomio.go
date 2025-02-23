//Autore: Francesco Corrado

package main

import(
    "fmt"
	"os"
	"strconv"
	"math"
	"bufio"
)

func main(){
	var x, n, somma int
	var y float64
	var s []string
	y = 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan(){
		numero := scanner.Text()
		s = append(s, numero)
	}
	x, _ = strconv.Atoi(s[len(s)-1])
	for i:=0; i < len(s)-1; i++{
		n, _ = strconv.Atoi(s[i])
		somma += n * int(math.Pow(float64(x),y))
		y += 1.0
	}
	fmt.Println(somma)
}