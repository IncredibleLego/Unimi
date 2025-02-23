//Autore: Francesco Corrado

package main

import(
	"fmt"
	"strconv"
	"os"
)

func Push(st []float64, x float64) []float64{
	st = append(st, x)
	return st
}

func Pop(st []float64) ([]float64, float64, bool){
	var testa float64
	if len(st) > 1{
		l := len(st)
		testa = st[l-1]
		st = st[:l-1]
		return st, testa, false
	}else{
		return st, 0, true
	}
}

func main(){
	var stack []float64
	var s string
	var n1, n2, out float64
	var err2 bool
	for{
		fmt.Print("Next? (+, -, *, /, q o un numero) ")
		fmt.Scan(&s)
		if s == "q"{
			os.Exit(0)
		}
		val, err := strconv.ParseFloat(s, 64)
		if err == nil{
			stack = Push(stack, val)
			fmt.Println(stack)
		}else{
			stack, n1, err2 = Pop(stack)
			if err2 == true{
				fmt.Println("not enough data")
				continue
			}
			stack, n2, _ = Pop(stack)
			if s == "+"{
				out = n1 + n2
			}else if s == "-"{
				out = n1 - n2
			}else if s == "*"{
				out = n1 * n2
			}else if s == "/"{
				out = n1 / n2
			}
			if s =="+" || s == "-" || s == "*" || s == "/"{
				stack = Push(stack, out)
				fmt.Println(stack)
			}
		}
	}
}

