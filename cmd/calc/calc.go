package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mbesseux/calc"
)

func main() {

	if len(os.Args) < 4 {
		fmt.Printf("usage:\n\t%s <operator> a b\n", os.Args[0])
		os.Exit(1)
	}
	a, _ := strconv.Atoi(os.Args[1])
	op := os.Args[2]
	b, _ := strconv.Atoi(os.Args[3])

	switch op {
	case "+":
		fmt.Println(calc.Add(a, b))
	case "-":
		fmt.Println(calc.Sub(a, b))
	case "*":
		fmt.Println(calc.Mult(a, b))
	case "/":
		ret, err := calc.Div(a, b)
		if err != nil {
			fmt.Println("Divide by 0!")
		} else {
			fmt.Println(ret)
		}

	default:
		fmt.Printf("%s operator not implemented", op)
	}

}
