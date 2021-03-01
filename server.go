package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func isFilePresent(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}
	return false
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	if r.URL.Path != "/ping" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	okFilePath := "/tmp/ok"
	if isFilePresent(okFilePath) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(http.StatusText(http.StatusOK)))
		return
	}
	w.WriteHeader(http.StatusServiceUnavailable)
	w.Write([]byte(http.StatusText(http.StatusServiceUnavailable)))
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(r)
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	if r.URL.Path != "/img" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	const Pixel = "\x47\x49\x46\x38\x39\x61\x01\x00\x01\x00\x80\xFF\x00\xFF\xFF\xFF\x00\x00\x00\x2C\x00\x00\x00\x00\x01\x00\x01\x00\x00\x02\x02\x44\x01\x00\x3B"
	w.Header().Set("Content-Type", "image/gif")
	w.Write([]byte(Pixel))
	fmt.Fprint(w)
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/img", imgHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
