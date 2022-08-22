package main

import "fmt"

func reverseString(str string) string {
	// Преобразуе строку в массив рун
	strRune := []rune(str)
	// Итерируеся по массиву рун и обмениваем симметричные элементы
	for pos, _ := range strRune {
		// Обмен
		buf := strRune[pos]
		strRune[pos] = strRune[len(strRune)-pos-1]
		strRune[len(strRune)-pos-1] = buf

		// Условие остановки если достигнута середина строки
		if pos == (len(strRune)-1)/2 {
			break
		}
	}
	return string(strRune)
}

func main() {
	// Пример عربى
	fmt.Print("Enter string: ")
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		return
	}
	fmt.Printf("Reversed: %s", reverseString(s))

}
