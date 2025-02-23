//Autore: Francesco Corrado

package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main(){
	scanner:= bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan(){
		riga:= scanner.Text()
		fmt.Println("Letta riga <" + riga + ">, di lunghezza " + strconv.Itoa(len(riga)) + " bytes")
		// CTRL + D per finire l'input
	}
}

