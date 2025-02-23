package main

import(
	"fmt"
	"math"
	"os"
	"strconv"
	"bufio"
	"strings"
)

type Wifi struct{
	ssid      	string
	channel   	int
	frequency 	int
	signal_dBm  int
}
/*
che, se i valori passati come parametri rispettano le seguenti condizioni:
	- ssid non vuoto
	- channel
		- tra 1 e 14 (compresi) SE la frequency è tra 2412 e 2484 (compresi)
		- OPPURE tra 7 e 196 SE la frequency è tra 5035 e 5980
	- frequency tra 2412 e 2484 OPPURE tra 5035 e 5980 (compresi estremi)
	- signal_dBm minore di -20 (meno venti)
	crea un'istanza di Wifi opportunamente inizializzata e la restituisce insieme a true, 
	altrimenti restituisce una struttura "zero-value" e false
*/
func NewWifi(ssid string, channel int, frequency int, signal_dBm int) (Wifi,bool){
	var wi Wifi
	wi.ssid = ""
	wi.channel = 0
	wi.frequency = 0
	wi.signal_dBm = 0
	if ssid != ""{
		if ((channel > 0 && channel < 15) && (frequency > 2411 && frequency < 2485)) || ((channel > 6 && channel < 197) && (frequency > 5034 && frequency < 5981)){
			if signal_dBm < -20{
				wi.ssid = ssid
				wi.channel = channel
				wi.frequency = frequency
				wi.signal_dBm = signal_dBm
				return wi, true
			}
		}
	}
	return wi, false
}
/*
che costruisce un'istanza della struttura Wifi a partire da una stringa nel formato CSV (in cui i dati sono 
	separati da virgole, vedere il file 'wifi.csv'), stessi vincoli della funzione definita sopra.
	Il formato è esattamente come nel file, non occorre fare controlli sul formato, ma i dati potrebbero essere 
	non accettabili (ad es. un numero di canale non coerente con la frequenza o l'intestazione del CSV).
*/
func NewWifiDaStringa(line string) (Wifi,bool){
	valori := strings.Split(line, ",")
	ssid := valori[0]
	channel, _ := strconv.Atoi(valori[1])
	frequency, _ := strconv.Atoi(valori[2])
	signal_dBm, _ := strconv.Atoi(valori[3])
	wi, check := NewWifi(ssid, channel, frequency, signal_dBm)
	return wi, check
}
/*
che restituisca una stringa rappresentativa dei valori della struttura, nella forma:
		{ssid, channel, frequency, signal_dBm, signalW}
	Attenzione che c'è un valore in più rispetto ai dati della struct, va aggiunto un valore calcolato: 
	la potenza del segnale in Watt. Il formato di uscita del valore signalW è quello "naturale" del float64 
	(formato "%v").
*/
func (wi Wifi) String() string{
	ssid := wi.ssid
	channel := wi.channel
	frequency := wi.frequency
	signal_dBm := wi.signal_dBm
	Watt := ConvertiDBinWatt(signal_dBm)
	output := fmt.Sprintf("{%s, %d, %d, %d, %v}\n", ssid, channel, frequency, signal_dBm, Watt)
//	output := fmt.Sprintln("{", ssid, channel, frequency, signal_dBm, Watt, "}")
	return output
}
/*
che prende come parametro la potenza del segnale in dBm (decibel-milliwatts) e calcola la potenza in Watt, 
la formula di conversione è:
		Watt = (10^(potenza_in_dBm/10)) / 1000
		Nota: il simbolo ^ indica elevamento a potenza (10^2 è 10 alla seconda)
*/
func ConvertiDBinWatt(signal_dBm int) float64{
	Watt := math.Pow(10, float64(signal_dBm/10)) / 1000
	return Watt
}
/*
che restituisce l'indice della struct che rappresenta il wifi con  segnale (dBm) più potente dell'elenco 
passato come parametro; in funzione del parametro banda viene restituito il più potente fra i segnali a 2GHz 
(banda="2GHz"), fra quelli a 5GHz (banda="5GHz") o senza distinzione (banda = qualunque altro valore, compresa 
la stringa vuota)
*/

func PiuPotente(elenco []Wifi, banda string) int{
	return 0
}

func main(){
	var valori []Wifi
	nFile := os.Args[1]
	file, _ := os.Open(nFile)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for scanner.Scan(){
		riga := scanner.Text()
		wi, err := NewWifiDaStringa(riga)
		if err == true{
			valori = append(valori, wi)
		}
	}
	fmt.Println(valori)
}