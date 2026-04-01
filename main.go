package main

import (
	"fmt"
	"strconv"

	"github.com/mthaler/aparser/ast"
	"github.com/mthaler/aparser/expr"
)

func main() {
	var roots []float64

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

		r, err := strconv.ParseFloat(s, 64)

		roots = append(roots, r)
		if err != nil {
			panic(err)
		}
	}

	for _, r := range roots {
		fmt.Println(r)
	}
}
