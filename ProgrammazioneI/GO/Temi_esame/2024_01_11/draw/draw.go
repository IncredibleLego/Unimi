package main

import(
	"fmt"
	"os"
	"strconv"
)
/*che restituisce una stringa formata da k spazi bianchi
seguiti dal carattere c */

func DrawPoint(c byte, k int) string{
	var output string
	for i:=0; i < k; i++{
		output += " "
	}
	output += string(c)
	return output
}

/* che restituisce una stringa formata da k spazi bianchi
seguiti da l caratteri c */

func DrawSegment(c byte, k, l int) string{
	var output string
	for i:=0; i < k; i++{
		output += " "
	}
	for i:=0; i < l; i++{
		output += string(c)
	}
	return output
}

func main(){
	var l, n int
	var c []byte
	s := os.Args[1]
	l, _ = strconv.Atoi(s)
	s = os.Args[2]
	n, _ = strconv.Atoi(s)
	s = os.Args[3]
	c = []byte(s)
	fmt.Println("l, n, c", l, n, c[0])
	if l > 0 && n > 0{
		fmt.Println(DrawPoint(c[0], l))
		fmt.Println(DrawSegment(c[0], n, l))
	}
}
