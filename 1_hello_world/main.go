package main

import (
	"flag"
	"fmt"
	"net/http"
)

var port = flag.String("port", ":8080", "listen port")

func main() {
	flag.Parse()

	http.HandleFunc("/", HelloHandler)

	fmt.Printf("Listening on %s.\n", *port)
	if err := http.ListenAndServe(*port, nil); err != nil {
		fmt.Printf("http.ListenAndServe %s failed: %s\n", *port, err)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello")
}
