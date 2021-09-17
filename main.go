package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world!")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Number of bytes written: %d\n", n)
	})

	address := ":8080"

	_ = http.ListenAndServe(address, nil)

}
