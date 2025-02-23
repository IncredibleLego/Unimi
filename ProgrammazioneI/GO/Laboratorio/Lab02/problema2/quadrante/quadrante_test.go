package main

import (
	//"strings"
	//"log"
	"testing"
)

var prog = "./quadrante"
var diffwidth = 120

func TestI(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"27687 89",
		"I quadrante\n")
}

func TestII(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"-2 0",
		"II quadrante\n")
}

func TestIII(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"-12 -45",
		"III quadrante\n")
}

func TestIV(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"127 -367",
		"IV quadrante\n")
}
