package youtube2rss

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

var YoutubeDlPath = "youtube-dl"

type feedConfig struct {
	RetrieveUrl string `json:"RetrieveUrl"`
	RetrieveLimit int `json:"RetrieveLimit"`
	PublishUrl  string `json:"PublishUrl"`
	Description string `json:"Description"`
	Title       string `json:"Title"`
	AuthorName  string `json:"AuthorName"`
	AuthorEmail string `json:"AuthorEmail"`
	FileFormat  string `json:"FileFormat"`
	ValidMatch  string `json:"ValidMatch"`
	TargetDir   string `json:"TargetDir"`
}

var FeedConfig feedConfig

func init() {
	readFlags()


	fmt.Printf("Config: %+v", FeedConfig)
}

func readConfigFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	Check(err, "Failed to read json config file ")

	err = json.Unmarshal(data, &FeedConfig)
	Check(err, "Failed to unmarshal json config data")
}

func readFlags() {
	retrieveUrl := flag.String("retrieveUrl", "", "url to read from")
	retrieveLimit := flag.Int("retrieveLimit", 5, "maximum number of items to retrieve")
	publishUrl := flag.String("publishUrl", "", "url to publish to")
	description := flag.String("description", "", "feed description")
	title := flag.String("title", "", "feed title")
	authorName := flag.String("authorName", "", "feed author name")
	authorEmail := flag.String("authorEmail", "", "feed author email")
	fileFormat := flag.String("fileFormat", "", "file format")
	configFile := flag.String("config", "", "config file")
	validMatch := flag.String("validMatch", "", "regex to match youtube video titles to")
	targetDir := flag.String("targetDir", "", "directory to download files to")

	flag.Parse()

	// Parse file first, to enable overriding with other flags
	if len(*configFile) > 0 {
		readConfigFile(*configFile)
	}

	if len(*retrieveUrl) > 0 {
		FeedConfig.RetrieveUrl = *retrieveUrl
	}

	FeedConfig.RetrieveLimit = *retrieveLimit

	if len(*publishUrl) > 0 {
		FeedConfig.PublishUrl = *publishUrl
	}

	if len(*description) > 0 {
		FeedConfig.Description = *description
	}

	if len(*title) > 0 {
		FeedConfig.Title = *title
	}

	if len(*authorName) > 0 {
		FeedConfig.AuthorName = *authorName
	}

	if len(*authorName) > 0 {
		FeedConfig.AuthorEmail = *authorEmail
	}

	if len(*fileFormat) > 0 {
		FeedConfig.FileFormat = *fileFormat
	}

	if len(*validMatch) > 0 {
		FeedConfig.ValidMatch = *validMatch
	}

	if len(*targetDir) > 0 {
		FeedConfig.TargetDir = *targetDir
	}
}
