package main

import (
	"fmt"
)

//Дан массив целых чисел, необходимо посчитать, количество уникальных значений
func SolutionUniq(a []int) int {
	b := make(map[int]int)

	for _, v := range a {
		b[v] = 0
	}
	return len(b)
}

func main() {
	a := []int{3, 8, 9, 7, 6, 7, 8, 1}
	fmt.Println(SolutionUniq(a))
}
