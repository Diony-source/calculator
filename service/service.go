package service

import (
	"calculator/entities"
	"errors"

	"golang.org/x/text/cases"
)

func PerformOperation(op entities.Operation) (float64, error) {
	switch op.Action {
	case "+":
		return op.Number1 + op.Number2, nil
	case "-":
		return op.Number1 - op.Number2, nil
	case "*":
		return op.Number1 * op.Number2, nil
	case "/":
		if op.Number2 == 0 {
		return 0, errors.New("division by zero is not allowed")
		}
		return op.Number1 / op.Number2, nil
	default:
		return 0, errors.New("unsupported operation")
	}
}