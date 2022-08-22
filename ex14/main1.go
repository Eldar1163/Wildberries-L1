package main

import (
	"fmt"
)

func getVariableType(variable interface{}) {
	// Можео определить тип с помощью данного выражения
	switch variable.(type) {
	case int:
		fmt.Println("Variable is int")
	case string:
		fmt.Println("Variable is string")
	case bool:
		fmt.Println("Variable is bool")
	case chan int:
		fmt.Println("Variable is channel")
	case chan string:
		fmt.Println("Variable is channel")
	case chan bool:
		fmt.Println("Variable is channel")
	}
}

func main() {
	var a interface{} = 1
	var b interface{} = "1"
	var c interface{} = true
	var d interface{} = make(chan int)
	var e interface{} = make(chan string)
	var f interface{} = make(chan bool)

	getVariableType(a)
	getVariableType(b)
	getVariableType(c)
	getVariableType(d)
	getVariableType(e)
	getVariableType(f)
}
