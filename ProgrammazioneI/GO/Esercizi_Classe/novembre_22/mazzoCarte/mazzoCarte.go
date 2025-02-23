//Autore: Francesco Corrado

package main

import(
	"fmt"
	"math/rand"
	"time"
)

type Carta struct{
	seme int //0=cuori, 1=quadri ecc.
	valore int //0=A, 1=2, 2=3...9=10, 10=J, 11=Q, 12=K
}

func mazzo() []Carta{ //Popola un mazzo di carte completo
	var m []Carta
	for s:=0; s<4; s++{
		for v:=0; v<13; v++{
			var c Carta
			c.seme = s
			c.valore = v
			m = append(m, c)
		}
	}
	return m
}

func mescola (m []Carta, r *rand.Rand){
	var j int
	n := len(m)
	res := make([]Carta, n)
	usato := make([]bool, n)
	for i:=0; i<n; i++{
		for{
			j= r.Intn(n)
			if usato[j]{
				break
			}
		}
		res[i] = m[j]
		usato[j] = true
	}
	for i:=0; i<n; i++{
		m[i] = res[i]
	}
}
/* VERSIONE DUE
func mescola (m []Carta, r *rand.Rand){
	n := len(m)
	for i:=0; i<n; i++{
		j:= i + r.Intn(n-1)
		m[i], m[j] = m[j], m[i]
	}
} */

func stampaCarte(m []Carta){
	for _, c:=range m{
		fmt.Println(Carta(c))
	}
}

func main(){
	r:= rand.New(rand.NewSource(time.Now().UnixNano()))
	m:=mazzo()
	mescola(m, r)
	stampaCarte(m[:5])
}