//Autore: Francesco Corrado

package main

import(
	"fmt"
	"math"
)

func main(){
	var P1x, P1y, P2x, P2y, P3x, P3y int
	var sulLatoVerticale, sulLatoOrizzontale bool
	var distanza12 float64
	fmt.Print("P1: ")
	fmt.Scan(&P1x, &P1y)
	fmt.Print("P2: ")
	fmt.Scan(&P2x, &P2y)
	fmt.Print("P3: ")
	fmt.Scan(&P3x, &P3y)
	if P1y <= P2y{
		sulLatoVerticale = (P3x == P1x || P3x == P2x) && (P3y >= P1y && P3y <= P2y)
	}else{
		sulLatoVerticale = (P3x == P1x || P3x == P2x) && (P3y >= P2y && P3y <= P1y)
	}
	if P1x <= P2x{
		sulLatoOrizzontale = (P3y == P1y || P3y == P2y) && (P3x >= P1x && P3x <= P2x)
	}else{
		sulLatoOrizzontale = (P3y == P1y || P3y == P2y) && (P3x >= P2x && P3x <= P1x)
	}
	distanza12 = math.Sqrt(float64((P2x-P1x)*(P2x-P1x) + (P2y-P1y)*(P2y-P1y)))
	

	fmt.Println(sulLatoVerticale)
	fmt.Println(sulLatoOrizzontale)
	fmt.Println(distanza12)
}