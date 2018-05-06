package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"github.com/pkg/errors"
)

var TargetDir = "output"
var YoutubeDlPath = "C:\\Users\\Thom\\scoop\\apps\\youtube-dl\\current\\youtube-dl.exe"

//var FeedUrl       =    // woody

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
	readConfigFile()
	readFlags()

	fmt.Printf("Config: %+v", FeedConfig)
}

func readConfigFile() {
	data, err := ioutil.ReadFile("./config.json")
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
	flag.Parse()

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

func Check(e error, info string) {
	if e != nil {
		log.Fatal(errors.Wrap(e, info))
	}
}
