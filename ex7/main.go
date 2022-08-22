package main

import (
	"fmt"
	"sync"
)

var (
	key  = "test"
	val1 = 5
	val2 = 12
)

func go1(m map[string]int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock() // Блокируем доступ к разделяемому ресурсу
	fmt.Println("Goroutine 1 start writing")
	m[key] = val1
	fmt.Println("Goroutine 1 stop writing")
	mutex.Unlock() // Разблокируем доступ к разделяемому ресурсу
	wg.Done()      // Указываем, что горутина завешила работу
}

func go2(m map[string]int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock() // Блокируем доступ к разделяемому ресурсу
	fmt.Println("Goroutine 2 start writing")
	m[key] = val2
	fmt.Println("Goroutine 2 stop writing")
	mutex.Unlock() // Разблокируем доступ к разделяемому ресурсу
	wg.Done()      // Указываем, что горутина завешила работу
}

func main() {
	wg := sync.WaitGroup{} // Создадим WaitGroup для ожидания завершениия работы горутин
	wg.Add(2)              // Укажем, что будет 2 горутины

	/*Создадим мьютекс, которы позволит установить
	монопольный доступ к общей переменной только
	1ой горутине в один момент времени */
	var mutex sync.Mutex

	m := make(map[string]int)
	go go1(m, &wg, &mutex)
	go go2(m, &wg, &mutex)

	wg.Wait() // Ждем завершения работы горутин
}
