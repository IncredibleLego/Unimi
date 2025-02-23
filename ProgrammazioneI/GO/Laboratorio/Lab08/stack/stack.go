//Autore: Francesco Corrado

package main

import(
	"fmt"
	"os"
)

func Push(st []float64) []float64{
	var x float64
	var sli1, sli2 []float64
	fmt.Println("valore?")
	fmt.Scan(&x)
	sli1 = append(sli1, x)
	sli2 = st
	st = append(sli1, sli2...)
	return st
}

func Pop(st []float64) []float64{
	if len(st) > 0{
		fmt.Println("testa", st[0])
		st = st[1:]
	}else{
		fmt.Println("Lo stack è troppo piccolo per rimuovere un valore")
	}
	return st
}

func Top(st []float64){
	if len(st) > 0{
		fmt.Println("testa", st[0])
	}else{
		fmt.Println("Lo stack non ha elementi impossibile trovare la testa")
	}
}

func Empty(st []float64) bool{
	if len(st) == 0{
		return true
	}else{
		return false
	}
}

func main(){
	var s string
	var stack []float64
	for{
		fmt.Println("Operazione (push/pop/top/empty/quit)?")
		fmt.Scan(&s)
		switch s{
			case "push":
				stack = Push(stack)
				fmt.Println(stack)
			case "pop":
				stack = Pop(stack)
				fmt.Println(stack)
			case "top":
				Top(stack)
			case "empty":
				fmt.Println(Empty(stack))
			case "quit":
				os.Exit(0)
			default:
				fmt.Println(s, "non è un opzione valida, riprova")
		}
	}
}