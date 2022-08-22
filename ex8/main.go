package main

import (
	"fmt"
	"math"
)

// Получает значение заданного бита в числе
func getBit(num int64, bit int) int64 {
	if bit > 0 {
		return (num >> (bit - 1)) % 2
	} else {
		panic("Wrong bit position")
	}
}

// Устанавливает значение заданного бита в числе
func bitSetter(num int64, bit int, val bool) int64 {
	mask := int64(math.Pow(2, float64(bit-1))) // Выделяем нужный бит
	if getBit(num, bit) == 1 && !val {
		return num &^ mask // Сброс единицы в ноль
	} else if getBit(num, bit) == 0 && val {
		return mask | num // Установка нуля в единицу
	} else {
		return num // Только два случая выше меняют значение исходного числа
	}
}

func main() {
	// 5 = 101
	// То что должно получиться смотреть в комментариях и сравнить с выводом
	fmt.Println(bitSetter(5, 1, true))  // 5 = 101
	fmt.Println(bitSetter(5, 1, false)) // 4 = 100
	fmt.Println(bitSetter(5, 2, true))  // 7 = 111
	fmt.Println(bitSetter(5, 2, false)) // 5 = 101
	fmt.Println(bitSetter(5, 3, true))  // 5 = 101
	fmt.Println(bitSetter(5, 3, false)) // 1 = 001
}
