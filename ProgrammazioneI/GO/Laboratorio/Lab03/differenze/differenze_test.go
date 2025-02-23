package main

// questo andrà tarato, potrebbero uscire decimali in più, anche i newline da specificare

import (
	//"strings"
	//"log"
	"testing"
)

var prog = "./differenze"
var diffwidth = 120

func Test1(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"3726.23 283928.6 0",
		"280202.37\n")
}

func Test2(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"2.1 2 0",
		"-0.1\n")
}

func Test3(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"8.5 10.0 0",
		"1.5\n")
}

func Test4(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"78.89 10.09 5.7 0",
		"-68.8\n-4.39\n")
}

func Test5(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"34.9 123.7 238.3535 0",
		"88.8\n114.6535\n")
}

func Test6(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"5.5 11.0 20.2 40.2 0",
		"5.5\n9.2\n20.0\n")
}
