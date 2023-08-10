package main

import (
	"strconv"
)

func Compute(op1 int, op2 int, operator string) string {
	switch operator {
	case "*":
		return strconv.Itoa(op1 * op2)
	case "/":
		return strconv.Itoa(op1 / op2)
	case "-":
		return strconv.Itoa(op1 - op2)
	case "+":
		return strconv.Itoa(op1 + op2)
	}

	return "Operator not valid. Please input *,/,-, or +"
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
