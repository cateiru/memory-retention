package main

import (
	"fmt"
	"net/http"

	memory_retention "github.com/yuto51942/memory-retention"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	result, err := memory_retention.CreateKey(ip)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, result)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	if err := memory_retention.AddAnswer(key, "hoge"); err != nil {
		fmt.Fprintln(w, err)
		return
	}
}

func getAllHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	result, err := memory_retention.GetAnswer(key)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, result)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/get", getAllHandler)

	http.ListenAndServe(":8080", nil)
}
