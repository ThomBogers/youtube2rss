package main

import (
	"log"
	"io/ioutil"
	"encoding/json"
)

var TargetDir = "output"
var YoutubeDlPath = "C:\\Users\\Thom\\scoop\\apps\\youtube-dl\\current\\youtube-dl.exe"
//var FeedUrl       =    // woody

type feedConfig struct{
	RetreiveUrl string `json:"RetreiveUrl"`
	PublishUrl string `json:"PublishUrl"`
	Description string `json:"Description"`
	Title string `json:"Title"`
	AuthorName string `json:"AuthorName"`
	AuthorEmail string `json:"AuthorEmail"`
}

var FeedConfig feedConfig

func init() {

	data, err := ioutil.ReadFile("./config.json")
	Check(err, "Failed to read json config file ")

	err = json.Unmarshal(data, &FeedConfig)
	Check(err, "Failed to unmarshal json config data")

}

func Check(e error, info string) {
	if e != nil {
		log.Fatal("Encoutered error: ", e, "(", info, ")")
	}
}
