package main

import (
	"fmt"
	"math/big"
)

/*
В данной задаче использован пакет math/big, который реализует длинную арифметику.
Все четыре арифметические функции, расположенные ниже возвращают в качестве результата строку,
представляющую результат арифметической операции в десятичной системе счисления
*/

// Сложение двух длинных чисел
func addition(num1 *big.Int, num2 *big.Int) string {
	res := new(big.Int)
	return res.Add(num1, num2).Text(10)
}

// Разность двух длинных чисел
func subtraction(num1 *big.Int, num2 *big.Int) string {
	res := new(big.Int)
	return res.Sub(num1, num2).Text(10)
}

// Произведение двух длинных чисел
func multiplication(num1 *big.Int, num2 *big.Int) string {
	res := new(big.Int)
	return res.Mul(num1, num2).Text(10)
}

// Частное двух длинных чисел (без остатка)
func division(num1 *big.Int, num2 *big.Int) string {
	res := new(big.Int)
	return res.Div(num1, num2).Text(10)
}

func main() {
	// Создание длинных чисел
	n1 := new(big.Int)
	n2 := new(big.Int)

	// Инициализация длинных чисел
	n1.SetString("10000000000000000000000000000000000000000000000012", 10)
	n2.SetString("10", 10)

	// Вывод на экран n1 и n2
	fmt.Printf("n1 = %s\n", n1.Text(10))
	fmt.Printf("n2 = %s\n", n2.Text(10))

	// Выполнение арифеметичсеких операций и вывод результатов
	fmt.Printf("n1 + n2 = %s\n", addition(n1, n2))
	fmt.Printf("n1 - n2 = %s\n", subtraction(n1, n2))
	fmt.Printf("n1 * n2 = %s\n", multiplication(n1, n2))
	fmt.Printf("n1 / n2 = %s", division(n1, n2))
}
