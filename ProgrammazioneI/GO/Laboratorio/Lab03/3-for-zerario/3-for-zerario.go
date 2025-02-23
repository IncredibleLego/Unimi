package main
import (
    "fmt"
    "math/rand"
//    "time"
)
func main() {
    /*
    Programma che genera numeri casuali tra 1 e 20 fino a che non genera uno 0, quindi stampa
	quanti numeri sono stati generati prima di ottenere lo 0 e termina
    */
    //rand.Seed(time.Now().Unix())
    const K = 20
    var n int

    c := 0
    for {
        n = rand.Intn(K+1)
        fmt.Print(n, " ")
        if n == 0 {
            break
        }
        c++
    }
    fmt.Println(c)
}
