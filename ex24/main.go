package main

import (
	"fmt"
	"math"
)

// Структура, представляющая точку

type Point struct {
	x float64
	y float64
}

// Функция расчета расстояния между двумя точками
func (p1 *Point) distance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2.0) + math.Pow(p2.y-p1.y, 2.0))
}

// Функция печати точки
func (p1 *Point) print() string {
	return fmt.Sprintf("{x = %f; y = %f}", p1.x, p1.y)
}

// Конструктор точки
func createPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}
func main() {
	// Создаем 2 точки
	p1 := createPoint(1, 5)
	p2 := createPoint(2, 5)

	// Выводим точки в консоль
	fmt.Printf("p1 = %s\n", p1.print())
	fmt.Printf("p2 = %s\n", p2.print())

	// Выводим расстояние между точками
	fmt.Printf("Distance between p1 and p2 = %f", p1.distance(p2))
}
