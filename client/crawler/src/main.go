package main

import (
	"crawler/client/crawler/src/client"
	"crawler/client/crawler/src/curl"
	"crawler/client/crawler/src/log"
	"crawler/client/crawler/src/notice"
	"github.com/PuerkitoBio/goquery"
	"time"
)

func main() {
	// protoc *.proto --go_out=plugins=grpc:../server
	for {
		time.Sleep(time.Second * 10)
		go run()
	}
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
	list := notice.GetList(document)
	client.FilterInsert(list)
}
