package main

import (
	"time"
)

type RssItem struct{
	Path string `json:"Path"`
	Link string `json:"Link"`
	VideoId   string `json:"VideoId"`
	Name string `json:"Name"`
	Date time.Time `json:"Date"`
}

