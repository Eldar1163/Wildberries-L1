package main

import (
	"fmt"
	"math/rand"
	"time"
)

func goroutine2(done chan bool) {
	fmt.Println("Goroutine was started")
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	for {
		//Периодически проверяем с помощью оператора select пришло, ли что-то в канал, если пришло, то завершаем горутину
		select {
		case <-done:
			fmt.Println("Goroutine was stopped")
			return
		default:
			fmt.Println(rnd.Intn(65))
			time.Sleep(400 * time.Millisecond)
		}
	}
}

func main() {
	ch := make(chan bool)
	go goroutine2(ch)

	time.Sleep(5 * time.Second)
	ch <- true // Отправляем в канал данные, чтобы горутина завершилась

	_, err := fmt.Scanln()
	if err != nil {
		return
	}
}
