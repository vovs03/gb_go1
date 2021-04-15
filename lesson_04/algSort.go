package main

// https://ideone.com/L63gQS#comments
// https://play.golang.org/p/Hcxh9pHS703 - дубликат
import (
	"fmt"
	"sort"
	"time"
)

func main() {
	p := fmt.Println
	s := []int{5, 24, 6, 30, 1, 4} // unsorted
	p(s)
	p("---")

	time.Sleep(2 * time.Second)
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	p("Отсортировано по убыванию: ", s)
	// p(s)

	p("---")

	time.Sleep(2 * time.Second)
	sort.Sort(sort.IntSlice(s))
	p("Отсортировано по возрастанию: ", s)
	//p(s)
}
