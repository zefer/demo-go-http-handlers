package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/zefer/demo_go_http_handlers/tweets"
)

var (
	tweetsClient = tweets.Client{}
	port         = flag.String("port", ":8080", "listen port")
)

func main() {
	flag.Parse()

	http.HandleFunc("/", TweetsHandler)

	fmt.Printf("Listening on %s.\n", *port)
	if err := http.ListenAndServe(*port, nil); err != nil {
		fmt.Printf("http.ListenAndServe %s failed: %s\n", *port, err)
		return
	}
}

func TweetsHandler(w http.ResponseWriter, r *http.Request) {
	tweets, err := tweetsClient.Fetch()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	for i, t := range tweets {
		fmt.Fprintf(w, "%d. %s (@%s)\n", i, t.Message, t.User)
	}
}
