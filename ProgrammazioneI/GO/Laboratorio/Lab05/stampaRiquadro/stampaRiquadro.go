//Autore: Francesco Corrado

package main

import "fmt"

func horizLine(ch byte, n int) string {
	var s, a string
	for i := 0; i <= n; i++ {
		a += string(ch)
	}
	s = " " + a + " "
	return s
}

func main() {
	var s, m string
	m = ""
	for {
		fmt.Scan(&s)
		if s == "stop" {
			break
		}
		if len(s) > len(m) {
			m = s
		}
	}
	fmt.Println(horizLine('-', len(m)))
	fmt.Println("|", m, "|")
	fmt.Println(horizLine('-', len(m)))
}
