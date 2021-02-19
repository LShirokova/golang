package main

import (
	"fmt"
	"strconv"
	"strings"
)

//
func SolutionBinaryGap(N int) int {
	i64 := int64(N)                  // Преобразую int в int64
	bin := strconv.FormatInt(i64, 2) // Преобразую N в двоичный код

	arr := strings.Split(bin, "") // Преобразую бинарный код в массив
	var i int                     // Счетчик
	var zeroLength int            // Сюда объем самой длинной последовательности

	// Прохожусь по масиву
	// и если элемент массива равен 0, то засчитываю его в последовательности,
	// иначе обнуляю счетчик
	for _, v := range arr {
		if v == "0" {
			i++
			if zeroLength < i {
				zeroLength = i
			}
		} else {
			i = 0
		}
	}

	return zeroLength
}

func main() {
	var example int = 1221 // <<< сюда подставляем число для проверки
	i64 := int64(example)
	fmt.Println(strconv.FormatInt(i64, 2), "Самая длинная последовательность: ", SolutionBinaryGap(example))
}
