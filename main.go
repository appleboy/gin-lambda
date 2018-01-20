package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
)

func main() {
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", hello)
	log.Fatal(gateway.ListenAndServe(addr, nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World from Go")
}
