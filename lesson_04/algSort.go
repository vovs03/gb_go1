package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arr := [4]int64{}

	scaner := bufio.NewScanner(os.Stdin)
	for i := len(arr) - 1; i >= 0; i-- {
		if scaner.Scan() {
			n, err := strconv.ParseInt(scaner.Text(), 10, 64)
			if err != nil {
				log.Println(err)
			}
			arr[i] = n
		}
	}

	//TODO sort
	fmt.Println(arr)
}
