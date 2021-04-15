package main

import "fmt"

func main() {
	// числа для приёма из консоли
	var a, b float32
	p := fmt.Println
	s := fmt.Scanln

	p("Введите первое число: ")
	s(&a)

	p("Введите второе число: ")
	s(&b)

	fmt.Printf("Сумма чисел: %f\n", a+b)

}
