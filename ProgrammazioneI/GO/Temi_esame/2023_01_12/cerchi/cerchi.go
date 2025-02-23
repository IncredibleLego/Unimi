package main

import(
	"fmt"
	"strings"
//	"io"
)

type Cerchio struct{
	nome string
	x, y, raggio float64
}

func newCircle(descr string) Cerchio{ //uno 0.5 0 2.5
	var c Cerchio
	_ = c
	array := strings.Split(descr, " ")
	fmt.Println(array)
	return c
}
/*
func distanzaPunti(x1,y1,x2,y2 float64) float64{

}

func tocca(cerc1, cerc2 Cerchio) bool{

}

func maggiore(cerc1, cerc2 Cerchio) bool{

}
*/
func main(){
	for{
		var s string
		var prec, att Cerchio
		prec = Cerchio{"", 0, 0, 0}
		_ = prec
		_ = att
		for{
			fmt.Scan(&s)
			fmt.Println("s: ", s)
		/*	if io.EOF == true{
				break
			} */

			att = newCircle(s)

		}

		//Bisogna leggere il nome, passare alle funzioni e salvare in una var il precedente
	}
}