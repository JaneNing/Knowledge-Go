package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println(reflect.TypeOf(x).Kind())
	fmt.Println(reflect.ValueOf(x))

	fmt.Println(reflect.TypeOf(x).Kind() == reflect.Float64)
	fmt.Println(reflect.TypeOf(x).Kind() == reflect.Float32)

	valueOf := reflect.ValueOf(x)
	fmt.Println(valueOf.CanSet())
	//valueOf.SetFloat(3.3) error
	newValueOf := reflect.ValueOf(&x)
	fmt.Println(newValueOf.CanSet())
	newValueOf = newValueOf.Elem()
	fmt.Println(newValueOf.CanSet())
	newValueOf.SetFloat(3.3)
	fmt.Println(newValueOf)
	fmt.Println(x)
}
