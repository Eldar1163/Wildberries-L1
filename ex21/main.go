package main

import "fmt"

/*
В данной задаче реализуется паттерн адаптер на примере температурного датчика, реализация которого сокрыта от нас.
Данный датчик умеет выдавать температуру только в градусах фаренгейта, а в нашей системе требуется вывод
в градусах цельсия. Для того, чтобы это сделать, была реализована структура AdapterSensor, которая хранит в себе
указатель датчик температуры и выполняет функцию преобразования одного интерфейса взаимодействия в нужный нам. По сути
происходит оборачивание одной структуры другою.
*/

// Интерфейс, к которому необходимо произвести адаптацию

type TempSensor interface {
	getCurrentTemp() float64
}

// Температурный сенсор, его реализация сокрыта, умеет только выдавать температуру в градусах фаренгейта

type ClosedTempSensor struct {
	tempInFahrenheit float64
}

// Интерфейс, который нужно адаптировать, выдает температуру в градусах фаренгейта
func (cts *ClosedTempSensor) getTemperature() float64 {
	return cts.tempInFahrenheit
}

// Структура адаптера, содержит в себе ссылку на температурный сенсор, интерфейс взаимодействия
// которого нужно адаптировать

type AdapterSensor struct {
	cts *ClosedTempSensor
}

// Функция реализующая адаптацию к необходимому интерфейсу взаимодействия
// (переводит градусы фаренгейта в градусы цельсия)
func (adapterSensor AdapterSensor) getCurrentTemp() float64 {
	return (adapterSensor.cts.tempInFahrenheit - 32.0) * 5.0 / 9.0
}

// Инициализация температурного датчика
var cts = ClosedTempSensor{tempInFahrenheit: 132}

func main() {
	// Создания экземпляра температурного датчика
	var adapterSensor TempSensor = AdapterSensor{cts: &cts}

	// Вывод температуры в требуемом формате
	fmt.Printf("Current temperature in celsius: %f", adapterSensor.getCurrentTemp())

}
