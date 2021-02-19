package main

import (
	"fmt"
)

// массив a, состоящий из n целых чисел
// Задача: сдвинуть массив a, k раз
func SolutionRotate(a []int, k int) []int {
	for i := 0; i < k; i++ {
		l := len(a)
		newArr := make([]int, l)
		newArr[0] = a[l-1]
		slice := a[0 : l-1]
		copy(newArr[1:], slice[:])
		a = newArr
	}

	return a
}

func main() {
	a := []int{3, 8, 9, 7, 6, 2}
	k := 2
	fmt.Println(SolutionRotate(a, k))
}
