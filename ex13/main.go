package main

import "fmt"

func swapNumbers(num1 *int, num2 *int) {
	// Здесь сумма num1 и num2
	*num1 += *num2
	// Здесь num1 + num2 - num2, то есть здесь num1
	*num2 = *num1 - *num2
	// Здесь num1 + num2 - num1, то есть здесь num2
	*num1 -= *num2
}

func main() {
	n1 := 2
	n2 := 5

	fmt.Printf("num1 = %d; num2 = %d\n", n1, n2)
	swapNumbers(&n1, &n2)
	fmt.Printf("num1 = %d; num2 = %d", n1, n2)
}
