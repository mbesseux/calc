package calc

import (
	"errors"
	"strconv"
)

// Add ajoute deux entiers
func Add(a int, b int) int {
	return a + b
}

// Sub soustraie deux entiers
func Sub(a int, b int) int {
	return a - b
}

// Mult multiplie deux entiers
func Mult(a int, b int) int {
	return a * b
}

// Div soustraie deux entiers
func Div(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Divide by 0")
	}
	return a / b, nil
}

// Operate parse!
func Operate(as, ops, bs string) (int, error) {

	a, _ := strconv.Atoi(as)
	b, _ := strconv.Atoi(bs)

	switch ops {
	case "+":
		return Add(a, b), nil
	case "-":
		return Sub(a, b), nil
	case "*":
		return Mult(a, b), nil
	case "/":
		return Div(a, b)

	default:
		return 0, errors.New("Operator not implemented")
	}
}
