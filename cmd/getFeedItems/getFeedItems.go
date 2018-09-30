package main

import (
	"fmt"
	"github.com/SlyMarbo/rss"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	. "github.com/ThomBogers/youtube2rss"
)

func main() {
	fmt.Print("Starting rss download\n")

	feed, err := rss.Fetch(FeedConfig.RetreiveUrl)
	Check(err, "Failed to fetch rss feed")

	fmt.Printf("Got rss feed, %d items\n", len(feed.Items))

	for _, item := range feed.Items {

		if !validTitle(item) {
			fmt.Printf("Skipping item %s\n", item.Title)
			continue
		}

		re := regexp.MustCompile(".*v=(.*)")

		if len(re.FindStringSubmatch(item.Link)) == 0 {
			log.Fatal("Could not parse videoId from link", item.Link)
		}

		videoId := re.FindStringSubmatch(item.Link)[1]
		outputFile := fmt.Sprintf("%s/%s.%s", TargetDir, videoId, FeedConfig.FileFormat)

		if existingFile(outputFile) {
			fmt.Printf("Existing item %s\n", item.Title)
			//When we encounter a file we have already downloaded we are done
			break
		}

		fmt.Printf("Got item: %s\n", item.Title)
		downloadFile(item.Link, videoId, outputFile)

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

func downloadFile(Link string, VideoId string, Path string) {
	fmt.Printf("Starting download for: %s to: %s\n", Link, Path)

	var cmd = exec.Command(YoutubeDlPath, "--print-json", "--extract-audio", "--audio-format", FeedConfig.FileFormat, "--audio-quality", "9", "--output", Path, Link)
	//cmd.Args = args

	stdout, err := cmd.StdoutPipe()
	Check(err, "Setting up Stderr pipe failed")

	err = cmd.Start()
	Check(err, "Starting command failed")

	data, err := ioutil.ReadAll(stdout)
	Check(err, "Sluping input failed")

	err = ioutil.WriteFile(fmt.Sprintf("%s/%s.json", TargetDir, VideoId), []byte(data), 0644)
	Check(err, "Failed to write yson to disk")

	err = cmd.Wait()
	Check(err, "Waiting for command failed")

	fmt.Printf("Finished download for %s\n", Link)
}
