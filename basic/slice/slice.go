package main

import "fmt"

func main() {
	slice := []int64{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	newSlice = append(newSlice, 60)

	fmt.Println(slice[3]) //60

	//slice := []int64{10, 20, 30, 40, 50}
	//newSlice := slice[2:3]
	//newSlice = append(newSlice, 60)
	//fmt.Println(slice[3]) //60
}
