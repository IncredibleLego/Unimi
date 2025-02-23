package main
import "fmt"
func main() {
    /*
	 Programma che dato un numero n int in input stampa su una riga n asterischi
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}
