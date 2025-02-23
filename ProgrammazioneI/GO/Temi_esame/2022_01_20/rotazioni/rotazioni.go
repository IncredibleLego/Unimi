package main

import(
    "fmt"
)

type Point struct{
    x, y float64
}

type Rectangle struct{
    pLL, pUR Point //Lower Left, Upper Right
}
/*
che data una coppia di coordinate x,y, restituisce il punto in quella posizione
*/
func NewPoint(x, y float64) Point{
	var p Point
	p.x = x
	p.y = y
	return p
}
/*
  che dati due vertici opposti di un rettangolo (P1 e P3 o P2 e P4, indifferentemente),
  restituisce il rettangolo corrispondente (vedi definizione di Rectangle), cioè in cui i due campi
  sono il vertice in basso a sinistra e quello il alto a destra (indicati con P1 e P3 nella figura sotto)

Nota: indipendentemente dai valori delle coordinate di due vertici opposti di un rettangolo (P1 e P3 o P2 e P4), P1 è dato da (min(x1,y1), min(y1,y2) e P3 da (max(x1,x1), max(y1,y2)

              P4 _____________ P3
                |             |
                |             |
                |_____________|
              P1               P2
*/

func NewRectangle(p1, p2 Point) Rectangle{
	var r Rectangle
	if p1.x < p2.x && p1.y < p2.y{
		r.pLL = p1
		r.pUR = p2
	}else{
		var pL, pU Point
		pL = NewPoint(p1.x, p2.y)
		pU = NewPoint(p2.x, p1.y)
	/*	pL.x = p1.x
		pL.y = p2.y
		pU.x = p2.x
		pU.y = p1.y */
		r.pLL = pL
		r.pUR = pU
	}
	return r
}
/*
- una funzione Rotate(r *Rectangle, dir byte)
che ruota il rettangolo r intorno al suo vertice pLL di 90 gradi in senso orario se dir è 'R' e in senso antiorario se dir è 'L'.
Per esempio, ruotando il rettangolo qui sopra in direzione 'L', il rettangolo diventa quello della figura qui sotto (P1 del vecchio rettangolo corrisponde a P2 del nuovo).
Per direzioni diverse da 'L' e 'R' la funzione non fa niente.

              P4 _____________ P2
                |             |
                |             |
                |_____________|
              P1               P3

      P1 _______ P4
        |       |
        |       |
        |       |
        |       |
        |_______|
      P3         P2                 

*/

func Rotate(r *Rectangle, dir byte){
	p1 := r.pLL
	p2 := r.pUR
	if dir == 'R'{
		pL := NewPoint(p1.x, p1.y - (p2.x - p1.x))
		pU := NewPoint(p1.x + (p2.y - p1.y), p1.y)
		r.pLL = pL
		r.pUR = pU
	}else if dir == 'L'{
		pL := NewPoint(p1.x - (p2.y - p1.y), p1.x)
		pU := NewPoint(p1.x, p2.y - p1.y)
		r.pLL = pL
		r.pUR = pU
	}
}
/*
he restituisce una riga di descrizione del rettangolo nel seguente formato: ((P1.x,P1.y) (P3.x,P3.y))
  (cioè ad es. "((2.1,3) (7,13.5))"
*/

func (re Rectangle) String() string{
    var output string
	p1 := re.pLL
	p2 := re.pUR
	output = fmt.Sprintf("((%3f,%3f) (%3f,%3f))", p1.x, p1.y, p2.x, p2.y)
	return output
}

func main(){
	p1 := NewPoint(10.2, 6)
	p2 := NewPoint(10.2, 20.4)
	r := NewRectangle(p1, p2)
	fmt.Println(r.String())
} 