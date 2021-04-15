package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv" // Преобразование строк в другие типы
	"strings" // ввод строк
)

// Весь этот листинг - абсолютный набросок, а не код!!!
func main() {
	return inputSide1()
	return inputSide2()
	return result()
}

func inputSide1() {

	var p = fmt.Println

	p("Введите длину стороны прямоугольника - ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	side1, err := strconv.ParseFloat(input, 64)

	logInformer()
	p("Вы ввели ", side1)
}

func logInformer() {
	if err != nil {
		p("")
		p("Нужно вводить числовые значения")
		log.Fatal(err)
	}
}

func result() {
	var Res = side1 * side2
	p(Res)
}
