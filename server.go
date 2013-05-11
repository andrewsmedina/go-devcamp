package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "devcamp 2013")
	})
	http.ListenAndServe("127.0.0.1:3333", nil)
}
