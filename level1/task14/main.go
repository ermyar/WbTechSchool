package main

import "fmt"

func deductType(value interface{}) string {
	var result string
	switch value.(type) {
	case int:
		result = "int"
	case string:
		result = "string"
	case bool:
		result = "bool"
	case chan int:
		result = "chan of int"
	case chan bool:
		result = "chan of bool"
	case chan string:
		result = "chan of string"
	case chan interface{}:
		result = "chan of any"
	default:
		result = "unknown type"
	}
	return result
}

func main() {
	fmt.Println(deductType(123))
	fmt.Println(deductType("hello"))
	fmt.Println(deductType(false))
	fmt.Println(deductType(make(chan int)))
	fmt.Println(deductType(make([]int, 1, 5)))
}
