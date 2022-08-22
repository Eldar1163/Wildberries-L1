package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func goroutine3(ctx context.Context) {
	fmt.Println("Goroutine was started")
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	for {
		//Если из контекста пришел сигнал завершение, то завершаем горутину
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine was stopped")
			return
		default:
			fmt.Println(rnd.Intn(65))
			time.Sleep(400 * time.Millisecond)
		}
	}
}

func main() {
	ctxt, cancel := context.WithCancel(context.Background()) // Используем контекст, для завершения работы горутины
	go goroutine3(ctxt)

	time.Sleep(5 * time.Second)
	cancel() // Завершаем горутину, при помощи данной функции

	_, err := fmt.Scanln()
	if err != nil {
		return
	}
}
