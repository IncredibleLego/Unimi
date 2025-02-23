package main
import (
	"fmt"
	"math/rand"
//	"time"
)
func main() {
    /*
    Programma che genera numeri casuali compresi fra 1 e 10 e li somma a una costante fino a che essa non supera 20
    */
    const TARGET = 20
    const MAX = 10
    //rand.Seed(time.Now().Unix())
    var n int

    somma := 0
    for somma < TARGET {
        n = rand.Intn(MAX) + 1
        fmt.Print(n, " ")
        somma += n
    }
    fmt.Print("\n",somma,"\n")
}
