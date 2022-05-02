package main

import "fmt"

func main() {
	getType("abc", true, float64(1), int64(1), float32(1))
}

func getType(list ...interface{}) {
	for _, item := range list {
		switch item.(type) {
		case bool:
			fmt.Println("bool")
		case float64:
			fmt.Println("float64")
		case int64:
			fmt.Println("int64")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("unknown")
		}
	}
}
