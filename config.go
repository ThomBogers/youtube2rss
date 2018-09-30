package youtube2rss

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

var TargetDir = "output"
var YoutubeDlPath = "youtube-dl"

type feedConfig struct {
	RetreiveUrl string `json:"RetreiveUrl"`
	PublishUrl  string `json:"PublishUrl"`
	Description string `json:"Description"`
	Title       string `json:"Title"`
	AuthorName  string `json:"AuthorName"`
	AuthorEmail string `json:"AuthorEmail"`
	FileFormat  string `json:"FileFormat"`
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
	publishUrl := flag.String("publishUrl", "", "url to publish to")
	description := flag.String("description", "", "feed description")
	title := flag.String("title", "", "feed title")
	authorName := flag.String("authorName", "", "feed author name")
	authorEmail := flag.String("authorEmail", "", "feed author email")
	fileFormat := flag.String("fileFormat", "", "file format")
	configFile := flag.String("config", "", "config file")
	flag.Parse()

	// Parse file first, to enable overriding with other flags
	if len(*configFile) > 0 {
		readConfigFile(*configFile)
	}

	if len(*retrieveUrl) > 0 {
		FeedConfig.RetreiveUrl = *retrieveUrl
	}

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
}
