//Autore: Francesco Corrado

package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main(){
	var p, q, s, somma float64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		riga := scanner.Text()
		str := strings.Split(riga, "#")
		p, _ = strconv.ParseFloat(str[0], 64)
		q, _ = strconv.ParseFloat(str[1], 64)
		s, _ = strconv.ParseFloat(str[2], 64)
		somma += (p * q - ((p*q * s*100)/100))
	}
	fmt.Println(somma)
}