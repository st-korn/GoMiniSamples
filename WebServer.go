package main

import (
	"fmt"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "<h1>Hello, world!</h1>")
}

func main() {
	http.HandleFunc("/",hello)
	http.ListenAndServe("localhost:4000",nil)
}