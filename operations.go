package main

import (
	"fmt"
)

func Compute(op1 int, op2 int, operator string) {
	switch operator {
	case "*":
		fmt.Println(op1 * op2)
	case "/":
		fmt.Println(op1 / op2)
	case "-":
		fmt.Println(op1 - op2)
	case "+":
		fmt.Println(op1 + op2)
	}
}

func GetArgs(equation string) (operand1, operand2, op string) {
	var op1, op2 string
	var operator string

	for i, v := range equation {
		if v < 48 || v > 57 {
			op1 = equation[0:i]
			op2 = equation[i+1:]
			operator = string(v)
			break
		}
	}

	return op1, op2, operator
}
