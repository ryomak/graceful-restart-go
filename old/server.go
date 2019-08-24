package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("This server is 1")
		fmt.Fprintf(w, "This server is 1")
	})
	lc := net.ListenConfig{}
	listener, err := lc.Listen(context.Background(), "tcp4", ":8080")
	if err != nil {
		panic(err)
	}
	http.Serve(listener, handler)
}
