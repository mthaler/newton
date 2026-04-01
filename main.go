package main

import (
	"fmt"
	"strconv"

	"github.com/mthaler/aparser"
	"github.com/mthaler/aparser/ast"
	"github.com/mthaler/aparser/expr"
)

func main() {
	var zeros []float64

	s := ""
	fmt.Println("Please enter a function:")
	fmt.Scanln(&s)
	fmt.Println(s)
	b := expr.CreateBuffer(s)
	e := expr.ArithmeticExpression()
	expr.Parse(e, b)
	a := ast.CreateAST(b.Code.Code)
	fmt.Printf("%+v\n", a)

	for {
		fmt.Println("Please enter a guess for the root or c to continue:")
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
