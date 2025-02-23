package main

import(
	"fmt"
	"strings"
	"strconv"
)

type BottigliaVino struct{
	nome string
    anno int
    gradi float32
    cl int
}

func CreaBottiglia(nome string, anno int, gradi float32, cl int) (BottigliaVino, bool){
	var bottle BottigliaVino
	if nome != "" && anno > 0 && gradi > 0 && cl > 0{
		bottle.nome = nome
		bottle.anno = anno
		bottle.gradi = gradi
		bottle.cl = cl
		return bottle, true
	}
	return bottle, false
}

func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool){
	valori := strings.Split(riga, ",")
	ann, _ := strconv.Atoi(valori[1])
	grad, _ := strconv.ParseInt(valori[2], 10, 32)
	gra := float32(grad)
	c, _ := strconv.Atoi(valori[3])
	output, boo := CreaBottiglia(valori[0], ann, gra, c)
	return output, boo
}

func main(){
	
}