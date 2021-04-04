package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv" // Преобразование строк в другие типы
	"strings" // ввод строк
)

func main() {

	// ======== return inputSide1()

	var p = fmt.Println

	p("Введите A: длину стороны прямоугольника - ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	side1, err := strconv.ParseFloat(input, 64)

	// logInformer()
	if err != nil {
		p("")
		p("Нужно вводить числовые значения")
		log.Fatal(err)
	}

	p("Вы ввели ", side1)

	// ======= return inputSide2()

	p("Введите B: длину стороны прямоугольника - ")

	reader2 := bufio.NewReader(os.Stdin)
	input2, err := reader2.ReadString('\n')

	input2 = strings.TrimSpace(input2)

	side2, err := strconv.ParseFloat(input2, 64)

	// ======== logInformer()
	if err != nil {
		p("")
		p("Нужно вводить числовые значения")
		log.Fatal(err)
	}

	p("Вы ввели ", side2)
	// return result()
	var Res = side1 * side2
	p("Площадь прямоугольника =", Res)

}
