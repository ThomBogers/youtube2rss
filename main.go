package main

import (
	"fmt"
	"github.com/SlyMarbo/rss"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"time"
	"encoding/json"
)

func main() {
	fmt.Print("Starting rss download\n")

	feed, err := rss.Fetch(FeedConfig.RetreiveUrl)
	Check(err, "Failed to fetch rss feed")

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
		outputFile := fmt.Sprintf("%s/%s.mp3", TargetDir, videoId)
		metaFile   := fmt.Sprintf("%s//%s.json", TargetDir, videoId)

		if existingFile(outputFile) {
			//When we encounter a file we have already downloaded we are done
			break
		}

		var rssitem = RssItem{
			outputFile,
			item.Link,
			videoId,
			item.Title,
			time.Now(),
		}

		fmt.Printf("Got item: %s\n", item.Title)

		itemJson, err := json.Marshal(rssitem)
		Check(err, "Failed to marshal meta data")

		err = ioutil.WriteFile(metaFile, itemJson, 0644)
		Check(err, "Failed to write meta data")

		downloadFile(rssitem)

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
	Check(err, "Failed to match item")

	return matched
}

func downloadFile(rssitem RssItem) {
	fmt.Printf("Starting download for: %s to: %s\n", rssitem.Link, rssitem.Path)

	var cmd = exec.Command(YoutubeDlPath, "--extract-audio", "--audio-format", "mp3", "--output", rssitem.Path, rssitem.Link)
	//cmd.Args = args

	stderr, err := cmd.StderrPipe()
	Check(err, "Setting up Stderr pipe failed")

	err = cmd.Start()
	Check(err, "Starting command failed")

	_, err = ioutil.ReadAll(stderr)
	Check(err, "Sluping input failed")

	err = cmd.Wait()
	Check(err, "Waiting for command failed")

	fmt.Printf("Finished download for %s\n", rssitem.Link)
}
