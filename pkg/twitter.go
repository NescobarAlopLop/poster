package pkg

import (
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type TwitterClient struct {
	Client *twitter.Client
}

func NewTwitterClient(consumerKey, consumerSecret string) *TwitterClient {
	config := &clientcredentials.Config{
		ClientID:     consumerKey,
		ClientSecret: consumerSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	httpClient := config.Client(oauth2.NoContext)
	client := twitter.NewClient(httpClient)

	return &TwitterClient{Client: client}
}

func (t *TwitterClient) GetPopularTweets(query string) ([]twitter.Tweet, error) {
	search, _, err := t.Client.Search.Tweets(&twitter.SearchTweetParams{
		Query: query,
		Count: 20,
	})

	if err != nil {
		return nil, err
	}

	return search.Statuses, nil
}

func (t *TwitterClient) PostTweet(text string) (*twitter.Tweet, error) {
	tweet, _, err := t.Client.Statuses.Update(text, nil)
	if err != nil {
		return nil, err
	}

	return tweet, nil
}
