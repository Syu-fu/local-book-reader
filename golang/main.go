package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", echoHello)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("error")
	}
}

func echoHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func EchoTest() {
	fmt.Printf("Test")
}
