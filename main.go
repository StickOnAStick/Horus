package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	protocols := http.Protocols{}
	protocols.SetHTTP1(true)
	protocols.SetHTTP2(false)

	helloWorld := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello world!\n")
	}

	helloHome := http.FileServer(http.Dir("/static"))

	http.HandleFunc("/hello", helloWorld)
	http.Handle("/", helloHome)

	log.Print("Listening on 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
