package notice

import (
	"crawler/client/crawler/src/log"
	"github.com/PuerkitoBio/goquery"
)

type Notice struct {
	Title      string
	CreateTime string
	JumpUrl    string
}

func GetList(document *goquery.Document) []Notice {
	lists := make([]Notice, 5)
	el := document.Find(".news_list_jigou > li")
	el.Each(func(i int, selection *goquery.Selection) {
		a := selection.Find(".title > a")
		url, _ := a.Attr("href")
		title := a.Text()
		time := selection.Find(".time").Text()
		notice := Notice{
			Title:      title,
			CreateTime: time,
			JumpUrl:    url,
		}
		log.Logger.Info(notice)
		lists = append(lists, notice)
	})
	return lists
}
