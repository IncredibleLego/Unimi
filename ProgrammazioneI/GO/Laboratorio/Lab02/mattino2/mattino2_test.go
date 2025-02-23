package main

import (
	//"strings"
	//"log"
	"testing"
)

var prog = "./mattino2"
var diffwidth = 120

func TestNoMattino1(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"0 7",
		"")
}

func TestNoMattino2(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"4 30",
		"")
}

func TestQuasiMattino(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"5 15",
		"")
}

func TestEdge1(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"5 30",
		"mattino\n")
}

func TestMattino(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"11 12",
		"mattino\n")
}

func TestEdge2(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"12 25",
		"mattino\n")
}

func TestEdge3(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"12 30",
		"")
}

func TestEdge4(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"13 0",
		"")
}

func TestNoMattino3(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"14 6",
		"")
}

func TestNoMattino4(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"19 30",
		"")
}
