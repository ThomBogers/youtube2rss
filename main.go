package main

import (
	"fmt"
	"github.com/SlyMarbo/rss"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
)

var YoutubeDlPath = "C:\\Users\\Thom\\scoop\\apps\\youtube-dl\\current\\youtube-dl.exe"

var FeedUrl = "https://www.youtube.com/feeds/videos.xml?channel_id=UCG9lNhVqk9luFLxBKDzuO9g" // 5secfilms for testing

/*
TODO
	- serialize meta data to id.meta  https://medium.com/@kpbird/golang-serialize-struct-using-gob-part-1-e927a6547c00
	- generate feed based on <id>.meta and <id.mp3> https://github.com/gorilla/feeds
*/

func main() {
	fmt.Print("Starting rss download\n")

	feed, err := rss.Fetch(FeedUrl)
	if err != nil {
		fmt.Print("Failed to fetch rss feed")
	}

	fmt.Print("Got rss feed\n")

	for _, item := range feed.Items {

		if !validTitle(item) {
			continue
		}

		re := regexp.MustCompile(".*v=(.*)")

		if len(re.FindStringSubmatch(item.Link)) == 0 {
			log.Fatal("Could not parse videoId from link", item.Link)
		}

		videoId := re.FindStringSubmatch(item.Link)[1]
		outputFile := fmt.Sprintf("./output/%s.mp3", videoId)

		if existingFile(outputFile) {
			//When we encounter a file we have already downloaded we are done
			break
		}


		fmt.Printf("Got item %s: %s %s\n", videoId, item.Title, item.Link)
		downloadFile(outputFile, item.Link)

		break
	}
}

func existingFile(f string) bool {
	_, err := os.Stat(f)

	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func validTitle(item *rss.Item) bool {
	matched, err := regexp.Match("PKA.*", []byte(item.Title))

	if err != nil {
		log.Fatal("Failed to match item")
	}

	return matched
}

func downloadFile(outputFile string, videoURL string) {
	fmt.Printf("Starting download for %s\n", videoURL)

	var cmd = exec.Command(YoutubeDlPath, "--extract-audio", "--audio-format", "mp3", "--output", outputFile, videoURL)
	//cmd.Args = args

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ := ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Finished download for %s\n", videoURL)
}
