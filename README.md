# Go net/http handler demo

This demonstrates a few simple concepts, useful when using & testing http
handlers. It shows different ways of writing a simple webservice with a single
handler which renders a list of tweets.

The `tweets` package shows how we might fetch and return tweet data, but it
returns fake data for demonstration purposes.

The directories prefixed `1_`, `2_`, etc show a progression of how we might
build up the code to create a working & testable http handler. The aim being to
gradually introduce the key concepts.
