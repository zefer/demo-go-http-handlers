package tweets

type Tweet struct {
	User    string
	Message string
}

type Client struct{}

// This is how the Fetch() method might look, except it would fetch real tweets
// from Twitter instead of returning placeholder data.
func (c Client) Fetch() ([]Tweet, error) {
	// TODO: fetch tweets from Twitter API & replace this placeholder data.
	tweets := []Tweet{
		{User: "zefer", Message: "Gorillas are big."},
		{User: "CodeCumbria", Message: "Food is nice."},
	}
	// TODO: return an error when fetching from Twitter fails.
	return tweets, nil
}
