package getFeedItems

import (
	"fmt"
	"github.com/SlyMarbo/rss"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"youtube2rss/config"
	"youtube2rss/util"
)

func main() {
	fmt.Print("Starting rss download\n")

	feed, err := rss.Fetch(config.FeedConfig.RetreiveUrl)
	util.Check(err, "Failed to fetch rss feed")

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
		outputFile := fmt.Sprintf("%s/%s.%s", config.TargetDir, videoId, config.FeedConfig.FileFormat)

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
	util.Check(err, "Failed to match item")

	return matched
}

func downloadFile(Link string, VideoId string, Path string) {
	fmt.Printf("Starting download for: %s to: %s\n", Link, Path)

	var cmd = exec.Command(config.YoutubeDlPath, "--print-json", "--extract-audio", "--audio-format", config.FeedConfig.FileFormat, "--audio-quality", "9", "--output", Path, Link)
	//cmd.Args = args

	stdout, err := cmd.StdoutPipe()
	util.Check(err, "Setting up Stderr pipe failed")

	err = cmd.Start()
	util.Check(err, "Starting command failed")

	data, err := ioutil.ReadAll(stdout)
	util.Check(err, "Sluping input failed")

	err = ioutil.WriteFile(fmt.Sprintf("%s/%s.json", config.TargetDir, VideoId), []byte(data), 0644)
	util.Check(err, "Failed to write yson to disk")

	err = cmd.Wait()
	util.Check(err, "Waiting for command failed")

	fmt.Printf("Finished download for %s\n", Link)
}
