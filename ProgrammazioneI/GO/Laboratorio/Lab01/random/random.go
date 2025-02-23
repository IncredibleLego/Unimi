//Autore: Francesco Corrado
package main
import "fmt"
import "math/rand"
import "time"
func main(){
	fmt.Println("Programma che stampa un numero casuale tra 0 e 100 ad ogni avvio")
	rand.Seed(time.Now().Unix())
	fmt.Println("Numero generato:", rand.Intn(100))
}