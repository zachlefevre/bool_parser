package boolParse

import (
	"fmt"
	"testing"
)

func TestB(t *testing.T) {
	inArr := []string{
		"T.$",
		"F.$",
		"T^F.$",
		"TvF.$",
		"(TvF) -> T v(T)  .   $",
		"TT.$",
		"(T).",
		"~~~T.",
	}
	var sArr []bool
	for _, e := range inArr {
		fmt.Println("input: ", e)
		get, s := CreateParser(e)
		inp := get()
		sArr = append(sArr, (B(&inp, get, s)))
	}
	fmt.Println(sArr)
}

func TestCreateParser(t *testing.T) {
	inArr := []string{
		"T.",
		"F.",
		"T^F.",
		"TvF.",
		"(TvF) -> T v(T)  .   ",
		"TT.",
		"T.$",
		"F.$",
		"T^F.$",
		"TvF.$",
		"(TvF) -> T v(T)  .   $",
		"TT.$",
		"(T).",
	}
	var pArr []func() byte
	var sArr [][]byte
	/* For each element of the array of examples, create a parser function */
	for _, e := range inArr {
		p, _ := CreateParser(e)
		pArr = append(pArr, p)
	}
	/* For each parser function, walk the enclosed string and append that string to sArr */
	for _, e := range pArr {
		var str []byte
		for r := e(); ; r = e() {
			str = append(str, r)
			if r == '$' {
				break
			}
		}
		sArr = append(sArr, str)
	}

	/* For each string parsed, print that string
	This is necessary over printing the sArr because each string actually a slice of bytes */
	for _, n := range sArr {
		fmt.Println(string(n))
	}
	/*
		For each element in the example array
			if the element ends in '$', then the corresponding sArr element should be identical*,
			if the element does not end in '$', then the corresponding sArr element should be equal to the
				example element + '$' */
}

func TestSemanticsOfL(t *testing.T) {
	input := "F."
	get, s := CreateParser(input)
	inp := get()
	B(&inp, get, s)
	fmt.Println(s.Peek())
}

func TestSemanticsOfAND(t *testing.T) {
	input := "T^T^T^T^F^T^T."
	get, s := CreateParser(input)
	inp := get()
	B(&inp, get, s)
	fmt.Println(s.Peek())
}

func TestSemanticsOfOR(t *testing.T) {
	input := "FvFvTvFvF."
	get, s := CreateParser(input)
	inp := get()
	B(&inp, get, s)
	fmt.Println(s.Peek())
}
