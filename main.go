package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/MihaiBogdanEugen/twittersearchgo"
	"github.com/namsral/flag"
	"github.com/sirupsen/logrus"
)

var (
	consumerKey      = flag.String("consumer-key", "", "Twitter API Consumer Key")
	consumerSecret   = flag.String("consumer-secret", "", "Twitter API Consumer Secret")
	outputFolderPath = flag.String("output-folder", "", "Output folder path")
	query            = flag.String("query", "", "Search query of 500 characters maximum, including operators")
	language         = flag.String("language", "en", "Restricts tweets to the given language, given by an ISO 639-1 code")
	resultType       = flag.String("result-type", "recent", "Specifies what type of search results you would prefer to receive")
	sinceID          = flag.Uint64("since-id", 0, "Returns results with an ID greater than (that is, more recent than) the specified ID")
	jsonLogging      = flag.Bool("json-logging", true, "Whether to log in JSON format or not")
	logLevel         = flag.String("log-level", "debug", "The log level (panic, fatal, error, warn, info, debug)")
	logger           = logrus.New()
)

func main() {
	flag.Parse()
	configLogging()

	client := configureClient(*consumerKey, *consumerSecret, *language, *resultType, *sinceID)
	logger.Info("Twitter Client created successfully")

	response, err := client.Search(*query)
	if err != nil {
		logger.WithError(err).Fatalf("cannot get query tweets by %s", *query)
	}

	if time.Now().Before(response.RateLimitReset) {
		logger.Infof("Rate limit resets at %v", response.RateLimitReset)
	}

	logger.Infof("Downloaded %d tweets for query term %s", len(response.Tweets), *query)

	if err := os.MkdirAll(*outputFolderPath, os.ModePerm); err != nil {
		logger.WithError(err).Fatalf("cannot create folder %s")
	}

	outputFilePath := filepath.Join(*outputFolderPath, fmt.Sprintf("%s.json", *query))

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		logger.WithError(err).Fatalf("cannot create output file %s", outputFilePath)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	for _, tweet := range response.Tweets {
		smallTweet := &SmallTweet{
			ID:             tweet.Id(),
			Text:           tweet.Text(),
			CreatedAt:      tweet.CreatedAt(),
			UserScreenName: tweet.User().ScreenName(),
			Language:       tweet.Language(),
		}

		tweetAsBytes, err := json.Marshal(smallTweet)
		if err != nil {
			logger.WithError(err).Fatalf("cannot json serialize object %v", smallTweet)
		}

		line := string(tweetAsBytes)
		if _, err := fmt.Fprintln(writer, string(tweetAsBytes)); err != nil {
			logger.WithError(err).Fatalf("cannot write line %s to output file %s", line, outputFilePath)
		}
	}

	if err := writer.Flush(); err != nil {
		logger.WithError(err).Fatalf("cannot flush writer of file %s", outputFilePath)
	}

	logger.Infof("Finished writing %d tweets to %s", len(response.Tweets), outputFilePath)
}

type SmallTweet struct {
	ID             uint64
	Text           string
	CreatedAt      time.Time
	UserScreenName string
	Language       string
}

func configureClient(twitterAPIConsumerKey string, twitterAPIConsumerSecret string, language string, resultType string, sinceID uint64) *twitterquerygo.SearchTwitterClient {
	client := twitterquerygo.NewClientUsingAppAuth(twitterAPIConsumerKey, twitterAPIConsumerSecret)
	client.SetLanguage(language)
	client.SetResultType(resultType)
	if sinceID > 0 {
		client.SetSinceID(sinceID)
	}
	return client
}

func configLogging() {
	if *jsonLogging {
		logger.Formatter = &logrus.JSONFormatter{}
	} else {
		logger.Formatter = &logrus.TextFormatter{}
	}

	level, err := logrus.ParseLevel(*logLevel)
	if err != nil {
		logger.Fatal("Invalid log level supplied", *logLevel)
	}

	logger.Out = os.Stdout
	logger.Level = level
}
