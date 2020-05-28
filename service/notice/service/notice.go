package noticeService

import (
	"context"
	"crawler/client/crawler/src/log"
	server "crawler/grpc/server"
	"crawler/service/mongo"
)

type NoticeServer struct {
}

func (receiver *NoticeServer) FilterInsert(ctx context.Context, in *server.FilterInsertRequest) (*server.FilterInsertReply, error) {
	Notices := in.Notices
	notices := make([]*server.Notice, 0)
	//	查询notices 然后过滤
	for _, value := range Notices {
		_, e := mongo.FindOneByTitle(value.Title)
		if e == nil {
			//	查询到了了就跳过
			continue
		}
		notices = append(notices, value)
	}
	if len(notices) != 0 {
		result, e := mongo.Insert(notices)
		if e != nil {
			panic(e)
		}
		log.Logger.Info("mongoId:", result.InsertedIDs)
	}
	return &server.FilterInsertReply{}, nil
}
