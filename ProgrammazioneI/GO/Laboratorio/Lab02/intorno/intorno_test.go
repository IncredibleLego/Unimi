package main

import (
	//"strings"
	//"log"
	"testing"
)

var prog = "./intorno"
var diffwidth = 120

func TestVicino1(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"0.000001 0.000005",
		"molto vicino all'origine\n")
}

func TestVicino2(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"0.000002 0.000003",
		"molto vicino all'origine\n")
}

func TestOrigine(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"0.0000 0.0000",
		"molto vicino all'origine\n")
}

func TestNonVicino1(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"0.001 0.000001",
		"non vicino all'origine\n")
}

func TestNonVicino2(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"3.6 2.7",
		"non vicino all'origine\n")
}

func TestEdge(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"0.00001 0.00001",
		"non vicino all'origine\n")
}
