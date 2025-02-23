//Autore: Francesco Corrado
package main
import "fmt"

func main(){
	var angolo1, angolo2 float64
	fmt.Println("Programma che chiede come input due angoli di un triangolo e calcola di conseguenza il terzo")
	fmt.Print("Inserire il primo angolo: ")
	fmt.Scan(&angolo1)
	fmt.Print("Inserire il secondo angolo: ")
	fmt.Scan(&angolo2)
	angolo3:=180-angolo1-angolo2
	fmt.Println("Dato l'angolo", angolo1, "e l'angolo", angolo2, ", il terzo angolo ha valore ", angolo3)
}