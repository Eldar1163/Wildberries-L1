package main

import (
	"fmt"
	"strings"
)

// Фунция проверяет строку на уникальность символов в ней без учета регистра
func uniqueChars(str string) bool {
	// Переводим все символы в нижний регистр, можно и в верхний разницы нет
	str = strings.ToLower(str)

	// Создаем map, за счет возможностей которой будем проверять уникальность
	m := make(map[int32]int)
	// Идем по символам строки
	for _, ch := range str {
		// Если ключ с текущим символом не найдем, то все хорошо, символ до этого не встречался
		if _, ok := m[ch]; !ok {
			m[ch] = 0
		} else {
			// Иначе, если ключ нашелся, то символ уже был и нужно вернуть false
			return false
		}
	}

	// Если дошли сюда, то все символы уникальные
	return true
}

func main() {
	fmt.Println(uniqueChars("abcd"))
	fmt.Println(uniqueChars("abCdefAaf"))
	fmt.Println(uniqueChars("aabcd"))
}
