package main

import (
	"fmt"
	"time"
)

// Реализация функции sleep при помощи функции after
func sleep1(t time.Duration) {
	_ = <-time.After(t)
}

//Реализация функции sleep при помощи функции NewTimer
func sleep2(t time.Duration) {
	timer := time.NewTimer(t)
	_ = <-timer.C
}

func main() {
	fmt.Println("Sleeping")

	sleep2(5 * time.Second)
	fmt.Println("Sleep1 success")

	sleep2(3 * time.Second)
	fmt.Printf("Sleep2 success")
}
