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

// We can call `tweets.Client` a `fetcher` because it satisfies this interface.
type fetcher interface {
	Fetch() ([]tweets.Tweet, error)
}

// Using the small 'fetcher' interface as the type of 'client' (rather than
// explicitly using `tweets.Client`) allows us to inject anything that has the
// Fetch method. This means we can easily test this handler's behaviour by
// providing a mock.
func TweetsHandler(c fetcher) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tweets, err := c.Fetch()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		for i, t := range tweets {
			fmt.Fprintf(w, "%d. %s (@%s)\n", i+1, t.Message, t.User)
		}
	})
}
