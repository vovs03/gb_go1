package main

import (
	"fmt"
	"os"
)

func main() {
	// числа для приёма из консоли + res(результат)
	var a, b, res float32
	// op | Operation
	var op string

	p := fmt.Println
	s := fmt.Scanln

	p("Введите первое число: ")
	s(&a)

	p("Введите второе число: ")
	s(&b)

	p("Введите арифметическую операцию (+, -, *, /): ")
	s(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	default:
		p("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат операции: %f\n", res)

}
