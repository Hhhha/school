package document

import (
	"crawler/client/crawler/src/curl"
	"crawler/client/crawler/src/log"
	"github.com/ddliu/go-httpclient"
)

func GetDocument(url string) *httpclient.Response {
	response, e := curl.Get(url, nil)
	if e != nil {
		log.Logger.Info(e)
	}
	return response
}
