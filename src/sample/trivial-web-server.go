package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Interactive Message</title>
	<style>
		body {
			font-family: Arial, sans-serif;
			background-color: #f0f0f0;
			text-align: center;
		}
		.container {
			margin-top: 100px;
			padding: 20px;
			background-color: #ffffff;
			border-radius: 10px;
			box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
			display: inline-block;
		}
		.message {
			color: #007bff;
			font-size: 24px;
		}
	</style>
	</head>
	<body>
		<div class="container">
			<span class="message">I am a GO application running inside a container - <span style="color: #ffc107;">v1.1</span>, some nigga have to level me up 👩‍🦱😅</span>
		</div>
	</body>
	</html>
	`
	fmt.Fprintf(w, html)
}

func main() {
	fmt.Println("Basic web server is starting on port 8080...")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
