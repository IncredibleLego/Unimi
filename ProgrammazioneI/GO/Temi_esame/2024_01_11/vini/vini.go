package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

type BottigliaVino struct{
	nome string
    anno int
    gradi float32
    cl int
}

func CreaBottiglia(nome string, anno int, gradi float32, cl int) (BottigliaVino, bool){
	var bottiglia BottigliaVino
	if nome != "" && anno > 0 && gradi > 0 && cl > 0{
		bottiglia.nome = nome
		bottiglia.anno = anno
		bottiglia.gradi = gradi
		bottiglia.cl = cl
		return bottiglia, true
	}else{
		bottiglia.nome = ""
		bottiglia.anno = 0
		bottiglia.gradi = 0
		bottiglia.cl = 0
		return bottiglia, false
	}
}

func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool){
	var bottiglia BottigliaVino
	var b bool
	bott := strings.Split(riga, ",")
	nome := bott[0]
	anno, _ := strconv.Atoi(bott[1])
	gr, _ := strconv.ParseFloat(bott[2], 32)
	gradi := float32(gr)
	cl, _ := strconv.Atoi(bott[3])
	bottiglia, b = CreaBottiglia(nome, anno, gradi, cl)
	return bottiglia, b
}

func (bottiglia BottigliaVino) String() string{
	var output string
	output = fmt.Sprintf("%s, %d, %g°, %dcl", bottiglia.nome, bottiglia.anno, bottiglia.gradi, bottiglia.cl)
	return output
}

func main(){
	var nBottiglie, tot int
	var gradMax, annoOld BottigliaVino
	nFile := os.Args[1]
	f, _ := os.Open(nFile)
	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		riga := scanner.Text()
		if riga != ""{
			bottiglia, _ := CreaBottigliaDaRiga(riga)
			fmt.Println(bottiglia.String())
			nBottiglie++
			tot += bottiglia.cl
			if bottiglia.gradi > gradMax.gradi{
				gradMax = bottiglia
			}
			if nBottiglie == 1{
				annoOld = bottiglia
			}
			if bottiglia.anno < annoOld.anno{
				annoOld = bottiglia
			}
		}
	}
	fmt.Println("n. bottiglie:", nBottiglie)
	fmt.Println("bottiglia con grado max:", gradMax.String())
	fmt.Println("bottiglia più vecchia:", annoOld.String())
	fmt.Printf("tot vino: %d cl\n", tot)
	f.Close()
}