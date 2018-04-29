package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"github.com/gorilla/feeds"
	"time"
	"log"
	"regexp"
	"encoding/json"
	"io/ioutil"
)

type ByFileDate []os.FileInfo

func (nf ByFileDate) Len() int      { return len(nf) }
func (nf ByFileDate) Swap(i, j int) { nf[i], nf[j] = nf[j], nf[i] }
func (nf ByFileDate) Less(i, j int) bool {
	// Use path names
	timeA := nf[i].ModTime()
	timeB := nf[j].ModTime()
	return timeA.Unix() < timeB.Unix()
}

func main() {
	var files []os.FileInfo
	var metas = make(map[string]RssItem)

	err := filepath.Walk(TargetDir, func(path string, info os.FileInfo, err error) error {
		fmt.Printf("p: %+v i: %+v\n", path, info)

		if filepath.Ext(path) == ".json" {

			data, err := ioutil.ReadFile(path)
			Check(err, "Failed to read json meta file ")

			var rssitem RssItem
			err = json.Unmarshal(data, &rssitem)
			Check(err, "Failed to unmarshal json meta data")

			metas[rssitem.VideoId]= rssitem

		} else if filepath.Ext(path) == ".mp3" {
			files = append(files, info)
		}
		return nil
	})

	Check(err, "Failed to walk path")

	sort.Sort(ByFileDate(files))
	sort.Reverse(ByFileDate(files))

	now := time.Now()
	feed := &feeds.Feed{
		Title:      FeedConfig.Title,
		Link:        &feeds.Link{Href: FeedConfig.RetreiveUrl},
		Description: FeedConfig.Description,
		Author:      &feeds.Author{Name: FeedConfig.AuthorName, Email: FeedConfig.AuthorEmail},
		Created:     now,
	}


	var feedItems []*feeds.Item

	for _, file := range files {
		fmt.Printf("Item: %s\n", file.Name())
		re := regexp.MustCompile("(.*).mp3")

		if len(re.FindStringSubmatch(file.Name())) == 0 {
			log.Fatal("Could not parse videoId from filename ", file.Name())
		}

		videoId := re.FindStringSubmatch(file.Name())[1]
		rssitem, exists := metas[videoId]

		if !exists {
			fmt.Println("Meta not found")
			continue
		}

		link        := &feeds.Link{ Href: fmt.Sprintf("%s/%s",FeedConfig.PublishUrl,file.Name())}
		description := fmt.Sprintf("youtube2rss feed item %s\n", rssitem.Name)
		item        := feeds.Item{
			Title: rssitem.Name,
			Link: link,
			Description: description,
			Created: rssitem.Date,
		}

		feedItems= append(feedItems, &item)
	}
	feed.Items = feedItems


	rss, err := feed.ToRss()
	Check(err, "Failed to create rss string")

	err = ioutil.WriteFile(fmt.Sprintf("%s/rss.xml", TargetDir), []byte(rss), 0644)
	Check(err, "Failed to write rss to disk")

	fmt.Printf("%s\n", rss)

}
