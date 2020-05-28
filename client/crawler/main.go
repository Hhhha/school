package crawler

import (
	"crawler/client/crawler/client"
	"crawler/client/crawler/curl"
	"crawler/client/crawler/log"
	"crawler/client/crawler/notice"
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
