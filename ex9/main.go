package main

import (
	"fmt"
	"sync"
)

func writer(arr []int, out chan int) {
	for _, val := range arr {
		out <- val
	}
	close(out) // Закрываем канал, чтобы оповестить получателя
}

func multiplyer(in chan int, out chan int) {
	for val := range in {
		out <- 2 * val
	}
	close(out) // Закрываем канал, чтобы оповестить получателя
}

func printer(in chan int, group *sync.WaitGroup) {
	for val := range in {
		fmt.Println(val)
	}
	group.Done() // Указываем, что горутина завершила работу
}

func main() {
	wg := sync.WaitGroup{} // Создание WaitGroup
	wg.Add(1)              // Говорим, что будем ждать одну горутину

	printMultCh := make(chan int)
	MultStdOutCh := make(chan int)

	go writer([]int{1, 2, 3, 4, 5}, printMultCh)
	go multiplyer(printMultCh, MultStdOutCh)
	go printer(MultStdOutCh, &wg)

	wg.Wait() // Ждем горутину, которая должна выполниться последней
}
