package main

import (
	"fmt"
)

var arr = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

// Находит группу для заданной температуры
func findGroup(temp float64) int {
	return int(temp) / 10 * 10
}

// Проходит по массиву температур и добавляет в map к нужной группе эту температуру
func makeGroups(arr []float64) {
	var res = make(map[int][]float64)
	for _, temp := range arr {
		res[findGroup(temp)] = append(res[findGroup(temp)], temp)
	}
	fmt.Println(res)
}

func main() {
	makeGroups(arr)
}
