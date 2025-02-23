//Autore: Francesco Corrado

package main

import(
    "fmt"
	"os"
	"strconv"
)

func main(){
	var s[] string //Crea l'array di stringhe dove mettere i valori di os.args
	var x, y int //Sono due variabili dove salvare il numero attuale x e il precedente y
	var err error //Variabile errore che servirà per verificare se s[i] è un umero
	s = os.Args[1:]
	for i:=0; i<len(s); i++{
		if i == 0{ //Se i = 0 è il primo numero quindi non va salvato nessun precedente 
				   //ne vanno fatti controlli
			x, err = strconv.Atoi(s[i])
			if err != nil{ //Se err è diverso da nil, allora il valore convertito non è un numero
				fmt.Println("Valore in posizione 0 non valido")
				break
			}
		}else{
			y = x //y = valore precendente
			x, err = strconv.Atoi(s[i]) //x = valore attuale
			if err != nil || i % 2 == 0 && x <= y || i % 2 != 0 && x >= y{
				//If con tutti i controlli dell'es: controlla che il numero sia effettivamente
				//un numero, che nel caso sia pari sia maggiore del precedente e nel caso sia 
				//dispari sia minore del precedente
				fmt.Println("Valore in posizione", i, "non valido.")
				break
			}
		}
		if i == len(s)-1{ //Se non sono stati rilevati errori, appena prima di finire il ciclo la
						  //Sequenza è valida
			fmt.Println("Sequenza valida.")
		}
	}
}