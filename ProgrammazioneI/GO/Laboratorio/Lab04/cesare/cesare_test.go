package main

import (
	//"strings"
	//"log"
	"testing"
)

var prog = "./cesare"
var diffwidth = 120

func Test0(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"0 abcdefg\n",
		"abcdefg\n")
}

func Test1(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"1 abcdefg\n",
		"bcdefgh\n")
}

func Test2(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"2 qwerty\n",
		"sygtva\n")
}

func Test3(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"3 thequickbrownfoxjumpsoverthelazydog\n",
		"wkhtxlfneurzqiramxpsvryhuwkhodcbgrj\n")
}

func Test4(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"40 thequickbrownfoxjumpsoverthelazydog\n",
		"hvseiwqypfckbtclxiadgcjsfhvszonmrcu\n")
}

func Test10(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"100 thequickbrownfoxjumpsoverthelazydog\n",
		"pdamqeygxnksjbktfqilokranpdahwvuzkc\n")
}
