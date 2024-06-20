package main

import "net/http"

func main() {
	server := http.ServeMux{}

	server.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "hello world"}`))
	})

	http.ListenAndServe(":8080", &server)
}
