package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	protocols := http.Protocols{}
	protocols.SetHTTP1(true)
	protocols.SetHTTP2(false)

	hanlde_homepage := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/static/index.html")
	}

	handle_helloWorld := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello world!\n")
	}

	handle_upload := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse input data..", http.StatusBadRequest)
			return
		}

		// Recieved message
		text := r.FormValue("text")
		fmt.Println(text)

		// Response
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `
			<!DOCTYPE HTML>
			<html>
			<head>
			</head>
			<body>
				<h2>Thank you!</h2>
				<p>You submitted: <strong>%s</strong><p>
				<a href="/">Back<a>
			</body>
			</html>
			`, text)

	}

	http.HandleFunc("/", hanlde_homepage)
	http.HandleFunc("/hello", handle_helloWorld)
	http.HandleFunc("/upload", handle_upload)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/static"))))

	log.Print("Listening on 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
