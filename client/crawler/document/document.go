package document

import (
	"crawler/client/crawler/curl"
	"crawler/client/crawler/log"
	"github.com/ddliu/go-httpclient"
)

func GetDocument(url string) *httpclient.Response {
	response, e := curl.Get(url, nil)
	if e != nil {
		log.Logger.Info(e)
	}
	return response
}
