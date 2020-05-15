package main

import (
	"crawler/client/crawler/src/curl"
	"crawler/client/crawler/src/log"
	"crawler/client/crawler/src/notice"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	run()
}

func run() {
	url := "http://www.hnswxy.com/danzhao/zhaoshenggonggao"
	response, e := curl.Get(url, nil)
	if e != nil {
		log.Logger.Error(e)
	}
	document, e := goquery.NewDocumentFromReader(response.Body)
	if e != nil {
		log.Logger.Error(e)
	}
	notice.GetList(document)
}
