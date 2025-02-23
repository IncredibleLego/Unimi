//Autore: Francesco Corrado

package main

import(
    "fmt"
	"math/rand"
	"strconv"
	"os"
	"os/exec"
	"time"
)

type Carta struct{
	valore, seme string
}

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

func carta(n int) (Carta, bool){
	var card Carta
	if n < 0 || n > 51{
		return card, false
	}
	v := n - (n/13)*13
	switch{
		case v == 0:
			card.valore = "A"
		case v == 10:
			card.valore = "J"
		case v == 11:
			card.valore = "Q"
		case v == 12:
			card.valore = "K"
		default:
			card.valore = fmt.Sprint(v+1)
	}
	switch{
		case n <= 12:
			card.seme = "\u2665"
		case n >= 13 && n <= 25:
			card.seme = "\u2666"
		case n >= 26 && n <= 38:
			card.seme = "\u2663"
		case n >= 39:
			card.seme = "\u2660"
	}
	return card, true
}

func getValoreBJ(card Carta) int{
	var output int
	if card.valore == "J" || card.valore == "Q" || card.valore == "K"{
		output = 10
	}else if card.valore == "A"{
		output = 11
	}else{
		output, _ = strconv.Atoi(card.valore)
	}
	return output
}

func mazzoPoker() []Carta{
	var mazzo []Carta
	for i:=0; i < 52; i++{
		card, _ := carta(i)
		mazzo = append(mazzo, card)
	}
	return mazzo
}

func mischia2mazzi(mazzo1 []Carta, mazzo2 []Carta) []Carta{
	mazzo := append(mazzo1, mazzo2...)
	for i:=0; i < 104; i++{
		r := rand.Intn(104)
		mazzo[i], mazzo[r] = mazzo[r], mazzo[i]
	}
	return mazzo
}

func preleva(mazzo []Carta)([]Carta, Carta){
	carta := mazzo[0]
	mazzo = mazzo[1:]
	return mazzo, carta
}

func stampaCarta(c Carta){
	var s string
	s = c.valore + c.seme
	fmt.Println(" ----- ")
	fmt.Println("|     |")
	fmt.Printf("|%4s |\n",s)
/*
	if len(s) == 3{
		fmt.Println("|", s, "|")
	}else{
		fmt.Println("|", s, " |")
	}
	*/
	fmt.Println("|     |")
	fmt.Println(" ----- ")
}

func stampaMano(mano []Carta){
	for i:=0; i < len(mano); i++{
		stampaCarta(mano[i])
	}
}

//Si potrebbe fare che finchè somma mano < sbando (21) asso = 11, altrimenti asso = 1
func calcolaMano(mano []Carta, giocatore bool) int{
	var somma int
	for i:=0; i < len(mano); i++{
		x := getValoreBJ(mano[i])
		if x == 11 && giocatore == true{
			for{
				fmt.Println("In mano hai un asso. Quanto vuoi che valga? (1 o 11)")
				fmt.Scan(&x)
				if x == 1 || x == 11{
					break
				}else{
					fmt.Println("Inserisci un valore valido (1 o 11)")
				}
			}
		}
		somma += x
	}
	return somma
}

func stampaTavolo(giocatore []Carta, mazziere []Carta){
	fmt.Println("Mano del mazziere: ")
	stampaMano(mazziere)
	fmt.Println("Tua mano: ")
	stampaMano(giocatore)
}

func vincitore(carteMazziere []Carta, cartePlayer []Carta, finaleMano int, cond string, vincitore string){
	calc := "calcolo del vincitore in corso"
	for i:=0; i < 4; i++{
		cancella()
		fmt.Println(calc)
		calc = calc + "."
		attendi(1)
	}
	cancella()
	fmt.Println("\t\tIL VINCITORE È\n")
	attendi(3)
	if cond == "playerSballato"{
		fmt.Println("IL BANCO: Vince il banco in quanto hai sballato")
	}else if cond == "mazziereSballato"{
		fmt.Println(vincitore, "! : Hai vinto dato che il banco ha sballato")
	}else{
		if calcolaMano(carteMazziere, false) >= finaleMano{
			fmt.Println("IL BANCO: ha vinto con un punteggio di", calcolaMano(carteMazziere, false), "punti contro i tuoi", finaleMano, "punti\n")
			if calcolaMano(carteMazziere, false) == finaleMano{
				fmt.Println("(in caso di parità vince il banco)\n")
			}
		}else{
			fmt.Println(vincitore, "! : Hai battutto il mazziere con un puteggio di", finaleMano, "contro il suo", calcolaMano(carteMazziere, false), "\n")
		}
	}
	fmt.Println("\t\tTavolo finale: ")
	stampaTavolo(cartePlayer, carteMazziere)
}

func Partita(){
	var nGiocatore string
	fmt.Println("Inserisci il nome con cui vuoi giocare ")
	fmt.Scan(&nGiocatore)
	var cartePlayer []Carta
	var carteMazziere []Carta
	var carta Carta
	var finaleMano int
	var fatto bool
	mazzo1 := mazzoPoker()
	mazzo2 := mazzoPoker()
	mazzo := mischia2mazzi(mazzo1, mazzo2)
	calc := "Creazione e mescolazione del mazzo in corso"
	for i:=0; i < 4; i++{
		cancella()
		fmt.Println(calc)
		calc = calc + "."
		attendi(1)
	}
	fmt.Println("Iniziamo! Ora il mazziere distribuirà due carte al giocatore e una al mazziere")
//	fmt.Println(mazzo)
	mazzo, carta = preleva(mazzo)
	cartePlayer = append(cartePlayer, carta)
	mazzo, carta = preleva(mazzo)
	cartePlayer = append(cartePlayer, carta)
	mazzo, carta = preleva(mazzo)
	carteMazziere = append(carteMazziere, carta)
	stampaTavolo(cartePlayer, carteMazziere)
	fmt.Println("La tua mano al momento vale", calcolaMano(cartePlayer, true))
	for{
		var scelta string
		fmt.Println("Seleziona l'azione (sto/chiamo)")
		fmt.Scan(&scelta)
		if scelta == "sto"{
			finaleMano = calcolaMano(cartePlayer, true)
			fmt.Println("Hai deciso di stare con la tua mano che vale", finaleMano)
			fmt.Println("Passaggio al gioco del mazziere")
			attendi(2)
			break
		}else if scelta == "chiamo"{
			fmt.Println("Il mazziere ti darà un altra carta")
			mazzo, carta = preleva(mazzo)
			cartePlayer = append(cartePlayer, carta)
			fmt.Println("Pesca in corso...")
			attendi(2)
			fmt.Println("Hai pescato la carta", carta)
			attendi(2)
			stampaTavolo(cartePlayer, carteMazziere)
			if calcolaMano(cartePlayer, true) > 21{
				fmt.Println("Hai sballato! la tua mano vale", calcolaMano(cartePlayer, true), "! (il limite è 21)")
				attendi(5)
				vincitore(carteMazziere, cartePlayer, finaleMano, "playerSballato", nGiocatore)
				finaleMano = -1
				break
			}
			fmt.Println("La tua mano al momento vale", calcolaMano(cartePlayer, true))
		}else{
			fmt.Println("Seleziona una delle opzioni disponibili")
		}
	}
	if finaleMano > -1{
		fmt.Println("Ora il mazziere pescherà fino a raggiungere o superare 17")
		for{
			mazzo, carta = preleva(mazzo)
			carteMazziere = append(carteMazziere, carta)
			stampaTavolo(cartePlayer, carteMazziere)
			fmt.Println("Il mazziere ha pescato", carta)
			if calcolaMano(carteMazziere, false) > 21{
				fmt.Println("Il mazziere ha sballato! la sua mano vale", calcolaMano(carteMazziere, false), "! (il limite è 21)")
				attendi(4)
				vincitore(carteMazziere, cartePlayer, finaleMano, "mazziereSballato", nGiocatore)
				fatto = false
				break
			}
			if calcolaMano(carteMazziere, false) > 16{
				fmt.Println("Il mazziere ha terminato il suo gioco")
				attendi(3)
				fatto = true
				break
			}
			fmt.Println("Pesca di un altra carta in corso")
			attendi(3)
		}
	}
	if finaleMano != -1 && fatto == true{
		vincitore(carteMazziere, cartePlayer, finaleMano, "", nGiocatore)
	}
}

func main(){
	var a string
	cancella()
	for{
		fmt.Println("\t\tBLACKJACK\n\nBenvenuto in questo simulatore di blackjack. \nInserisci PLAY per giocare e EXIT per terminare il programma")
		fmt.Scan(&a)
		if a == "PLAY"{
			fmt.Println("Perfetto! Iniziamo")
			attendi(2)
			cancella()
			break
		}else if a == "EXIT"{
			fmt.Println("Grazie per aver usato questo simulatore")
			os.Exit(0)
		}else{
			fmt.Println("Inserire una delle opzioni disponibili")
			attendi(2)
			cancella()
		}
	}
	for{
		Partita()
		for{
			fmt.Println("\nVuoi giocare ancora? (y/n)")
			fmt.Scan(&a)
			if a == "y"{
				fmt.Println("Perfetto, inizializzo una nuova partita")
				attendi(2)
				cancella()
				break
			}else if a == "n"{
				fmt.Println("Grazie per aver usato questo simulatore")
				os.Exit(0)
			}else{
				fmt.Println("selezionare una delle opzioni disponibili")
				attendi(2)
				cancella()
			}
		}
	}
}