//Autore: Francesco Corrado

package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<5; i++{
		fmt.Println(rand.Intn(11))
	}
}

r:= rand.New(rand.NewSource(time.now().UnixNano()))
	for i:=0; i<5; i++{
		fmt.Println(r.Intn(11))
	}