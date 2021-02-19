package main

import (
	"fmt"
	"strconv"
)

// Генератор квадратов натуральных чисел
// Выводит квадраты n-количества чисел начиная от числа start
func SolutionSquareGenerator(start int, n int) string {
	var results = ""
	for i := start; i <= n; i++ {
		results += strconv.Itoa(i*i) + " "
	}
	return results
}

func main() {
	fmt.Println(SolutionSquareGenerator(2, 4)) // 4 9 16
}
