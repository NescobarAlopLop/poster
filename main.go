package main

import (
	"fmt"
	"github.com/NescobarAlopLop/poster/pkg"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"os"
	"time"
)

func main() {
	twitterConsumerKey := os.Getenv("TWITTER_API_KEY")
	twitterConsumerSecret := os.Getenv("TWITTER_API_SECRET_KEY")

	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")

	openaiApiKey := os.Getenv("OPENAI_API_KEY")

	config := oauth1.NewConfig(twitterConsumerKey, twitterConsumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	twitterClient := pkg.NewTwitterClient(twitterConsumerKey, twitterConsumerSecret)
	openaiClient := pkg.NewClient(openaiApiKey)

	if len(os.Args) < 2 {
		fmt.Println("Using default query: 'bitcoin'")
		os.Args = append(os.Args, "bitcoin")
	} else {
		fmt.Println("Using query:", os.Args[1])
	}

	twitterQuery := os.Args[1]
	for {
		tweets, err := twitterClient.GetPopularTweets(twitterQuery)
		if err != nil {
			fmt.Println("Error fetching tweets:", err)
			continue
		}

		tweetTexts := make([]string, len(tweets))
		for i, tweet := range tweets {
			tweetTexts[i] = tweet.Text
		}

		summary, err := openaiClient.SummarizeTweets(tweetTexts)
		if err != nil {
			fmt.Println("Error summarizing tweets:", err)
			continue
		}

		tweet, _, err := client.Statuses.Update(summary, nil)
		if err != nil {
			fmt.Println("Error posting tweet:", err)
			continue
		}
		fmt.Printf("Tweet ID: %v\n", tweet.ID)

		time.Sleep(time.Hour)
	}
}
