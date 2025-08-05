package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello Argo - 05082025 - 1549</h1>"))
	})
	http.ListenAndServe(":8080", nil)
}
