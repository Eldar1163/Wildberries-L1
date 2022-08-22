package main

import (
	"fmt"
	"reflect"
)

func getType(variable interface{}) {
	// Можно определить тип с помощью функции пакета reflect
	switch reflect.TypeOf(variable).Kind() {
	case reflect.Int:
		fmt.Println("Variable is int")
	case reflect.String:
		fmt.Println("Variable is string")
	case reflect.Bool:
		fmt.Println("Variable is bool")
	case reflect.Chan:
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

	getType(a)
	getType(b)
	getType(c)
	getType(d)
	getType(e)
	getType(f)
}
