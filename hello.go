package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var port string

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World! This is a GO app.")
}

func main() {
	flag.StringVar(&port, "port", "5000", "HTTP Server port")
	flag.Parse()

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	http.HandleFunc("/", HelloHandler)

	fmt.Printf("Listening on %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
