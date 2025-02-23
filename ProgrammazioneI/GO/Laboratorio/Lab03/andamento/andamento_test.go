package main

// questo andrà tarato, potrebbero uscire decimali in più, anche i newline da specificare

import (
	//"strings"
	//"log"
	"testing"
)

var prog = "./andamento"
var diffwidth = 120

func Test1(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"2 4 7 3 9 1 5 -1",
		"++-+-+\nsomma: 31\n")
}

func Test2(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"1 2 3 4 5 6 7 8 -1",
		"+++++++\nsomma: 36\n")
}

func Test3(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"1000 100 50 30 20 10 -1",
		"-----\nsomma: 1210\n")
}

func Test4(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"0 10 5 32 20 100 -1",
		"+-+-+\nsomma: 167\n")
}

func Test5(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"10 20 30 50 100 1000 -1",
		"+++++\nsomma: 1210\n")
}
