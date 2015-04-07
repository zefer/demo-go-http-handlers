package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/zefer/demo_go_http_handlers/tweets"
)

var port = flag.String("port", ":8080", "listen port")

func main() {
	flag.Parse()

	client := tweets.Client{}

	http.Handle("/", TweetsHandler(client))

	fmt.Printf("Listening on %s.\n", *port)
	if err := http.ListenAndServe(*port, nil); err != nil {
		fmt.Printf("http.ListenAndServe %s failed: %s\n", *port, err)
		return
	}
}

// Wrapping the handler in a closure allows us to pass in our tweets client,
// thus avoiding the need for a global variable.
func TweetsHandler(c tweets.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tweets, err := c.Fetch()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		for i, t := range tweets {
			fmt.Fprintf(w, "%d. %s (@%s)\n", i, t.Message, t.User)
		}
	})
}
