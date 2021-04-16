package main

// https://ideone.com/X4DROj
//  Хотелось проверить в плейграунде, но запустил только локально.

import (
	"fmt"
	"strconv"
)

func main() {
	var s string

	fmt.Println("Введите порядковый номер числа из ряда Фибоначчи.")
	_, errScan := fmt.Scan(&s)
	n, err := strconv.ParseInt(s, 10, 0)
	if n <= 0 || errScan != nil || err != nil {
		fmt.Println("Число должно быть целым и положительным!")
		return
	}

	func() {
		f := FibonacciWithDefer()
		for i := 0; i < int(n)-1; i++ {
			f()
		}

		fmt.Println("Число Фибоначчи:", f())
	}()
}

func FibonacciWithDefer() func() int {
	a, b := 1, 1
	return func() int {
		defer func() {
			a = a + b
			a, b = b, a
		}()
		return a
	}
}
