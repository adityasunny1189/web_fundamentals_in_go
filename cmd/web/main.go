package main

import (
	"fmt"
	"net/http"

	"github.com/adityasunny1189/web_fundamentals_in_go/pkg/handler"
)

const PORT = ":5000"

func main() {
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/about", handler.AboutHandler)

	fmt.Printf("server running on Port: %v\n", PORT)
	_ = http.ListenAndServe(PORT, nil)
}
