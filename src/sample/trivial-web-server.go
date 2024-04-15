package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	const (
		Reset  = "\033[0m"
		Red    = "\033[31m"
		Green  = "\033[32m"
		Yellow = "\033[33m"
	)

	message := fmt.Sprintf("%sI am a GO application running inside Docker - %sV2.4.7%s, and Just now I have been upgraded%s",
		Green, Yellow, Green, Reset)

	fmt.Fprintf(w, message)
}

func main() {
	fmt.Println("Basic web server is starting on port 8080...")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
