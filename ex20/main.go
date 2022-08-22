package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Функция разворачивает порядок слов в строке
func reverseWords(str string) string {
	// Результирующая строка изначально пустая
	res := ""
	// Делим строку на слова по пробелам
	words := strings.Split(str, " ")

	// Читаем массив с конца и записываем результат в res
	for i := len(words) - 1; i >= 0; i-- {
		res += words[i]
		// Не забываем добавлять пробел, когда это нужно
		if i != 0 {
			res += " "
		}
	}
	return res
}

func main() {
	fmt.Print("Enter string: ")

	// Для чтения строки используем сканер
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	fmt.Printf("Reversed words order: %s", reverseWords(s))
}
