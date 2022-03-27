package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/post", Handler)
	http.ListenAndServe("localhost:8000", nil)
}

func Handler(writer http.ResponseWriter, request *http.Request) {
	name := 999
	fmt.Fprint(writer, "hello "+fmt.Sprintf("%d", name))
}
