package main

import (
	//"strings"
	//"log"
	"testing"
)

var prog = "./euclide"
var diffwidth = 120

func Test1(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"80 56",
		"8\n")
}

func Test2(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"790 40",
		"10\n")
}

func Test3(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"1357 987",
		"1\n")
}

func Test4(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"7328423874 73847384",
		"2\n")
}

func Test5(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"1111111260 34568",
		"4\n")
}
