package curl

import "github.com/ddliu/go-httpclient"

func Get(url string, param map[string]string) (*httpclient.Response, error) {
	response, e := httpclient.Get(url, param)
	return response, e
}
func Post(url string, param map[string]string) (*httpclient.Response, error) {
	response, e := httpclient.Post(url, param)
	return response, e
}
