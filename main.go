package main

import (
	"fmt"
	"strconv"

	"github.com/mthaler/aparser"
)

func main() {
	var zeros []float64

	for {
		s := ""
		fmt.Println("Please enter a guess for the zero or q to quit")
		fmt.Scanln(s)
		if s == "q" {
			break
		}

		z, err := strconv.ParseFloat(s, 64)

		zeros = append(zeros, z)
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("%+v\n", zeros)

	r, err := aparser.Eval("3 + 4")
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
