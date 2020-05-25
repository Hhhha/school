package notice

import (
	"crawler/client/crawler/src/log"
	grpc "crawler/grpc/server"
	"github.com/PuerkitoBio/goquery"
)

func GetList(document *goquery.Document) []*grpc.Notice {
	lists := make([]*grpc.Notice, 0)
	el := document.Find(".news_list_jigou > li")
	el.Each(func(i int, selection *goquery.Selection) {
		a := selection.Find(".title > a")
		url, _ := a.Attr("href")
		title := a.Text()
		time := selection.Find(".time").Text()
		notice := &grpc.Notice{
			Title:      title,
			CreateTime: time,
			JumpUrl:    url,
		}
		log.Logger.Info(notice)
		lists = append(lists, notice)
	})
	return lists
}
