//Autore: Francesco Corrado

package main

import(
	"fmt"
)

type Segmento struct{
	estremi byte
	interno byte
	orizzontale bool
	lunghezza int
}

func (se Segmento) String() string {
	var output string
	if se.lunghezza == 1{
		output = output + "*"
	}else if se.lunghezza == 0{
		output = ""
	}else{
		output = output + string(se.estremi)
		if se.orizzontale == false{
			output = output + "\n"
		}
		for x:=0; x < se.lunghezza-2; x++{
			output = output + string(se.interno)
			if se.orizzontale == false{
				output = output + "\n"
			}
		}
		output = output + string(se.estremi)
	}
	output = output + "\n"
	return output
}

func (se *Segmento) allunga(n int) string {
	var output string
	se.lunghezza += n
	output = se.String()
	return output
}

func main(){
	var s Segmento
	var x string
	var f bool
	var i int
	var b []byte
	fmt.Scan(&x)
	b = []byte(x)
	s.estremi = b[0]
	fmt.Scan(&x)
	b = []byte(x)
	s.interno = b[0]
	fmt.Scan(&f)
	s.orizzontale = f
	fmt.Scan(&i)
	s.lunghezza = i
	fmt.Print(s.String())
	fmt.Print(s.allunga(3))
	s.orizzontale = !s.orizzontale
	fmt.Print(s.String())
}