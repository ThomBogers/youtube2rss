package main

import (
	"encoding/json"
	"fmt"
	"github.com/eduncan911/podcast"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"time"
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
	var ymeta = make(map[string]YoutubeDlData_s)

	err := filepath.Walk(TargetDir, func(path string, info os.FileInfo, err error) error {

		if filepath.Ext(path) == ".json" {

			data, err := ioutil.ReadFile(path)
			Check(err, "Failed to read yson meta file ")

			var youtbeDlData YoutubeDlData_s
			err = json.Unmarshal(data, &youtbeDlData)
			Check(err, "Failed to unmarshal json meta data")

			ymeta[youtbeDlData.ID] = youtbeDlData

		} else if filepath.Ext(path) == fmt.Sprintf(".%s", FeedConfig.FileFormat) {
			files = append(files, info)
		}
		return nil
	})

	Check(err, "Failed to walk path")

	sort.Sort(ByFileDate(files))
	sort.Reverse(ByFileDate(files))

	now := time.Now()
	//feed := &feeds.Feed{
	//	Title:       FeedConfig.Title,
	//	Link:        &feeds.Link{Href: FeedConfig.RetreiveUrl},
	//	Description: FeedConfig.Description,
	//	Author:      &feeds.Author{Name: FeedConfig.AuthorName, Email: FeedConfig.AuthorEmail},
	//	Created:     now,
	//}

	p := podcast.New(
		FeedConfig.Title,
		FeedConfig.RetreiveUrl,
		FeedConfig.Description,
		&now,
		&now,
	)

	p.AddAuthor(FeedConfig.AuthorName, FeedConfig.AuthorEmail)
	p.AddCategory("Comedy", nil)
	p.AddImage(fmt.Sprintf("%s/image.png", FeedConfig.PublishUrl))

	regexString := fmt.Sprintf("(.*).%s", FeedConfig.FileFormat)
	re := regexp.MustCompile(regexString)

	for _, file := range files {
		fmt.Printf("Item: %s\n", file.Name())

		if len(re.FindStringSubmatch(file.Name())) == 0 {
			log.Fatal("Could not parse videoId from filename ", file.Name())
		}

		videoId := re.FindStringSubmatch(file.Name())[1]
		rssitem, exists := ymeta[videoId]

		if !exists {
			fmt.Println("Meta not found")
			continue
		}

		pubDate, err := time.Parse("20060102", rssitem.UploadDate)

		item := podcast.Item{
			Title:       rssitem.Title,
			Description: fmt.Sprintf("youtube2rss feed item %s\n", rssitem.Description),
			PubDate:     &pubDate,
			GUID:        rssitem.ID,
		}
		item.AddEnclosure(fmt.Sprintf("%s/data/%s", FeedConfig.PublishUrl, file.Name()), getType(FeedConfig.FileFormat), file.Size())

		_, err = p.AddItem(item)
		Check(err, "Failed to add item to feed")
	}

	rss := p.String()

	Check(err, "Failed to create rss string")

	err = ioutil.WriteFile(fmt.Sprintf("%s/rss.xml", TargetDir), []byte(rss), 0644)
	Check(err, "Failed to write rss to disk")

	fmt.Printf("%s\n", rss)

}

func getType(typeName string) podcast.EnclosureType {
	switch typeName {
	case "mp3":
		return podcast.MP3
	case "m4a":
		return podcast.M4A
	case "m4v":
		return podcast.M4V
	case "mp4":
		return podcast.MP4
	case "mov":
		return podcast.MOV
	case "pdf":
		return podcast.PDF
	case "epub":
		return podcast.EPUB
	default:
		return podcast.M4A
	}
}
