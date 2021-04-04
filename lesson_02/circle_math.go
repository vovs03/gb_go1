package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv" // Преобразование строк в другие типы
	"strings" // ввод строк
)

func main() {
	const Pi = 3.14159265358979323846264338327950288419716939937510582097494459
	// g = greeting
	var g = "Введите площадь круга"
	var p = fmt.Println

	p(g) //
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	circle, err := strconv.ParseFloat(input, 64)

	if err != nil {
		p("")
		p("Нужно вводить числовые значения")
		log.Fatal(err)
	}

	var D = 2 * math.Sqrt(circle/Pi)
	var l = Pi * D
	p("Диаметр круга = ", D)
	p("Длина окружности = ", l)
}
