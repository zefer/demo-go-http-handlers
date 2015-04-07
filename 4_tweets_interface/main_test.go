package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zefer/demo_go_http_handlers/tweets"
)

var (
	client  *mockTweetsClient
	handler http.Handler
	w       *httptest.ResponseRecorder
)

// 'Mock' the tweets client so we can test the handler's behaviour.
type mockTweetsClient struct{}

// 'Stub' the 'Fetch' method to return some fake tweets, we'll assert that this
// data is rendered correctly in the handler's response body.
func (c mockTweetsClient) Fetch() ([]tweets.Tweet, error) {
	return []tweets.Tweet{
		{User: "zefer", Message: "Monkeys stand for honesty"},
		{User: "CodeCumbria", Message: "Giraffes are insincere"},
		{User: "zefer", Message: "Orangutans are skeptical of changes in their cages"},
	}, nil
}

// Another 'Mock' of the tweets client used to test a scenario where it fails.
type mockFailingTweetsClient struct{}

// 'Stub' Fetch() to return an error, to simulate a failure.
func (c mockFailingTweetsClient) Fetch() ([]tweets.Tweet, error) {
	return []tweets.Tweet{}, errors.New("Something went wrong.")
}

// This interface allows us to re-use our test setup() method, as we can pass in
// any mock of the tweets client as long as it satisfies this interface.
type Fetcher interface {
	Fetch() ([]tweets.Tweet, error)
}

// Common test setup. Run the handler with a ResponseRecorder to capture the
// response, so we can assert it looks as expected.
func setup(client Fetcher) {
	w = httptest.NewRecorder()
	handler = TweetsHandler(client)
	req, _ := http.NewRequest("GET", "/", nil)
	handler.ServeHTTP(w, req)
}

// When the tweet client returns tweets.
func TestTweetsHandlerOK(t *testing.T) {
	setup(&mockTweetsClient{})

	expectedBody := "1. Monkeys stand for honesty (@zefer)\n"
	expectedBody += "2. Giraffes are insincere (@CodeCumbria)\n"
	expectedBody += "3. Orangutans are skeptical of changes in their cages (@zefer)\n"
	if w.Body.String() != expectedBody {
		t.Errorf("Body is `%s` expected `%s`", w.Body.String(), expectedBody)
	}

	if w.Code != http.StatusOK {
		t.Errorf("Code is %d expected %d", w.Code, http.StatusOK)
	}
}

// When the tweet client fails (returns an error).
func TestTweetsHandlerError(t *testing.T) {
	setup(&mockFailingTweetsClient{})

	if w.Body.String() != "" {
		t.Errorf("Expected empty body, got `%s`", w.Body.String())
	}

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Code is %d expected %d", w.Code, http.StatusInternalServerError)
	}
}
