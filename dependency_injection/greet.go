package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	greetingPrefix = "Hello, "
)

func Greet(writer io.Writer, name string) {
	// 1 - via Fprintf
	fmt.Fprintf(writer, greetingPrefix+name)
	// 2 - via inbuilt method
	// // Dereferenced as it is a struct
	// buff.WriteString(greetingPrefix + name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(
		http.ListenAndServe(
			":5001",
			http.HandlerFunc(MyGreeterHandler),
		),
	)
}
