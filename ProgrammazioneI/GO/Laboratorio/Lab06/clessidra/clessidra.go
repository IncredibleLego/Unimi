//Autore: Francesco Corrado

package main

import(
	"fmt"
	"strings"
	"os/exec"
	"os"
	"time"
)

// Funzione che fa il clear dello schermo
// importare "os/exec"
func cancella() {
	cmd := exec.Command("clear")
	cmd.Stdout=os.Stdout
	cmd.Run()
}

// Funzione che attende
// time.Sleep(time.Duration(n_seconds) * time.Second)
// (o time.Millisecond, time.Nanosecond....)
func attendi(n_seconds float64){
	time.Sleep(time.Duration(n_seconds) * time.Second)
}

// rigaClessidra restituisce la una singola riga di sabbia o di non sabbia (interna alla clessidra)
// accetta lunghezza, stato di sabbia o no, carattere da usare per la sabbia;
// usare le funzioni di strings (senza usare il for)
func rigaClessidra(length int, sand bool, sandChar byte) string {
	var output string
	if sand{
		output += strings.Repeat(string(sandChar), length)
	}else{
		output += strings.Repeat(" ", length)
	}
	return output
}

// funzione che disegna tutta la clessidra,
// per ogni riga, il cono e l'aria o la sabbia,
// accetta altezza, secondi passati dallo start, char per la sabbia
// l'altezza sara` il numero di secondi totali
// e sara` l'altezza di ciascun "cono"
// usate (altezza * 2 + 1) come base per il "cono"
func clessidra(height int, time int, char byte){
	var n, a int
	n = 0
	a = height*2 - 1
	for i:=0; i < height; i++{
		fmt.Print(rigaClessidra(n, false, ' '))
		fmt.Print("\\")
		if i < time{
			fmt.Print(rigaClessidra(a, true, ' '))
		}else{
			fmt.Print(rigaClessidra(a, true, '*'))
		}
		fmt.Print("/")
		n++
		a -= 2
		fmt.Println()
	}
	n = height-1
	a = 1
	for i:=height; i > 0; i--{
		fmt.Print(rigaClessidra(n, false, ' '))
		fmt.Print("/")
		if i > time{
			fmt.Print(rigaClessidra(a, true, ' '))
		}else{
			fmt.Print(rigaClessidra(a, true, '*'))
		}
		fmt.Print("\\")
		n--
		a += 2
		fmt.Println()
	}
}

func main(){
	var t int
	fmt.Scan(&t)
	cancella()
	for i:=0; i < t+1; i++{
		clessidra(t, i, '*')
		attendi(1)
		if i < t{
			cancella()
		}
	}
}