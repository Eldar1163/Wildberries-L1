package main

import (
	"fmt"
	"sync"
	"time"
)

// Счетчик
var counter int

func worker(workerId int, m *sync.Mutex, done chan struct{}) {
	// Бесконечный цикл инкрементиирования счетчика
	for {
		// Обработка сигнала завершения работы
		select {
		case _, ok := <-done:
			if !ok {
				return
			}
		default:
		}
		//Блокировка доступа других горутин к разделяемому ресурсу
		m.Lock()
		// Инкрементиирование счетчика
		counter++
		// Вывод значения счетчика с указанием какой воркер вывел
		fmt.Printf("Worker %d, counter = %d\n", workerId, counter)
		//Разблокировка доступа других горутин к разделяемому ресурсу
		m.Unlock()
		// Небольшая задержка для наглядности работы программы
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	// Вывод начального значения счетчика
	fmt.Printf("Main, counter initial value = %d\n", counter)

	// Канал для сигнала завершения работы горутин
	done := make(chan struct{})
	// Мьютекс для корректного доступа к разделяемым ресурсам (счетчику)
	var mutex = sync.Mutex{}

	// Запуск горутин-воркеров
	go worker(1, &mutex, done)
	go worker(2, &mutex, done)

	// Задержка
	time.Sleep(5 * time.Second)
	// Подача "сигнала" завершения работы горутинам
	close(done)

	// Вывод конечного значения счетчики
	fmt.Printf("Main, counter finish value = %d", counter)
}
