package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	fmt.Println("Starting test server on :3923")
	if err := http.ListenAndServe(":3923", nil); err != nil {
		fmt.Println("Error starting test server:", err)
	}
}
