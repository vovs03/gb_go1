package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	// числа для приёма из консоли + res(результат)
	var a, b, res float64
	// op = 'Operation'
	var op string

	p := fmt.Println
	s := fmt.Scanln

	p("Введите первое число: ")
	s(&a)
	if a != float64(a) {
		//break
		p("Число введено неверно")
		os.Exit(2)
	}

	p("Введите второе число: ")
	s(&b)

	p("Введите арифметическую операцию (+, -, *, /, sqrt): ")
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
	case "sqrt":
		res = math.Sqrt(a*a + b*b)
	default:
		p("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат операции: %f\n", res)

}
