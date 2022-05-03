package main

import (
	"flag"
	"fmt"
)

func main() {
	//fmt.Println(len(os.Args))
	//fmt.Println(os.Args[1:])
	//go run args.go test

	s := flag.String("content", "none", "")
	res := flag.Bool("result", false, "")
	num := flag.Int64("num", 1, "")
	flag.Parse()
	if s != nil {
		fmt.Println(*s)
	}
	if res != nil {
		fmt.Println(*res)
	}
	if num != nil {
		fmt.Println(*num)
	}
	//go run args.go -content="abc" -result=true -num=2
}
