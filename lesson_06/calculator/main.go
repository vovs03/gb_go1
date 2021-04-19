package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const helpMessage = "Формат вычисления: 'a / b'"

// структура калькулятора
type calcCalls struct {
	// Nums: a, b & res = result
	a, b, res float64
	// op - operation
	op string
}

// передача структуры в переменную `myCalc`
var myCalc calcCalls

func sum(a, b float64) (float64, error) {
	return a + b, nil
}

func subtract(a, b float64) (float64, error) {
	return a - b, nil
}

func multiply(a, b float64) (float64, error) {
	return a * b, nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Делить на ноль нельзя!")
	}
	return a / b, nil
}

/*
мапа соответствий названия действия функции-обработчику
(список методов действий калькулятора)
*/
var opMap = map[string]func(float64, float64) (float64, error){
	"+": sum,
	"-": subtract,
	"*": multiply,
	"/": divide,
}

func processOp(a float64, op string, b float64) (float64, error) {
	if opFunc, ok := opMap[op]; ok {
		return opFunc(a, b)
	}
	return 0, fmt.Errorf("Не соблюдён вормат ввода")
}

func main() {
	fmt.Println("Доступные операции калькулятора: +, -, *, /")
	scanner := bufio.NewScanner(os.Stdin)
	// Запуск цикла калькулятора( построчно вести рассчёты)
	for scanner.Scan() {
		input := scanner.Text()
		// токен, он же счётчик трёх аргументов формулы ( a op b )
		tokens := []string{}

		// strings.Split разбивает строку по заданному символу
		for _, token := range strings.Split(input, " ") {
			token = strings.TrimSpace(token) // отрезаем пробелы
			if token != "" {
				tokens = append(tokens, token)
			}
		}
		if len(tokens) != 3 {
			fmt.Println(helpMessage)
			continue
		}
		op := tokens[1]
		a, err1 := strconv.ParseFloat(tokens[0], 64)
		b, err2 := strconv.ParseFloat(tokens[2], 64)
		if err1 != nil || err2 != nil {
			fmt.Println(helpMessage)
			continue
		}
		res, err := processOp(a, op, b)
		if err != nil {
			fmt.Printf("%s: %v\n", input, err)
			continue
		}
		fmt.Println("Результат:", res)
	}
}
