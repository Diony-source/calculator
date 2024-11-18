package handlers

import (
	"calculator/entities"
	"calculator/service"
	"fmt"
)

func start() {
	var num1, num2 float64
	var action string

	fmt.Println("Enter the first number:")
	fmt.Scanln(&num1)

	fmt.Println("Enter the second number:")
	fmt.Scanln(&num2)

	fmt.Println("Enter the operation (+, -, *, /):")
	fmt.Scanln(&action)
	
	operation := entities.Operation{
		Number1: num1,
		Number2: num2,
		Action: action,
	}

	result, err := service.PerformOperation(operation)
	if err != nil {
		fmt.Println("Error:", err)
	}else {
		fmt.Println("The result is: %.2f\n", result)
	}
}