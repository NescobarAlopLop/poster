# Poster - Twitter Summarizer Bot (Go)

This project is a fun little bot that uses the Twitter API to fetch the most popular tweets based on a provided search query. It then uses the OpenAI GPT-3 API to generate a summary of these tweets and posts the summary back to Twitter as a new tweet.

## Current State

As of now, the bot is capable of:

- Fetching popular tweets from Twitter based on a search query.
- Extracting the text and any mentioned URLs from these tweets.
- Using the OpenAI API to generate a summary of the extracted text.

In the future, we plan to add:

- More sophisticated summarization capabilities.
- Better error handling and logging.
- More configuration options for the search query and summarization.

## How to Use

To use the bot, you will need to provide your Twitter API keys and OpenAI API key as environment variables. You can set these variables in your shell like so:

```bash
export TWITTER_CONSUMER_KEY=your_twitter_api_key
export TWITTER_CONSUMER_SECRET=your_twitter_api_secret_key
export TWITTER_ACCESS_TOKEN=your_twitter_access_token
export TWITTER_ACCESS_TOKEN_SECRET=your_twitter_access_token_secret
export OPENAI_API_KEY=your_openai_api_key
```

Then, you can run the bot with a search query like this:

```bash
go run main.go "your-twitter-query"
```

Replace `"your-twitter-query"` with the actual query you want to use. For example, if you wanted to search for tweets containing the word "golang", you would run:

```bash
go run main.go "golang"
```

## Disclaimer

This project is just for fun and is not intended for serious use. Please use responsibly and respect the terms of service of both Twitter and OpenAI.

---
All of the above and most of the code was generated by ChatGPT, I had little to do with the results. I just wanted to see what it can do.
