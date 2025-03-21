package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello Argo - 1701</h1>"))
	})
	http.ListenAndServe(":8080", nil)
}
