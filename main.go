package main

import (
	"fmt"
	"strconv"

	"github.com/mthaler/aparser"
)

func main() {
	var zeros []float64

	s := ""
	fmt.Println("Please enter a function:")

	for {
		fmt.Println("Please enter a guess for the zero or c to continue:")
		fmt.Scanln(&s)
		if s == "c" {
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
