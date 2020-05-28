package client

import (
	"context"
	"crawler/client/crawler/log"
	grpc2 "crawler/grpc/server"
	"crawler/service/notice/config"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn
var noticeServiceClient grpc2.NoticeServiceClient

func init() {
	connect()
}
func connect() {
	conn, e := grpc.Dial("127.0.0.1:"+config.PORT, grpc.WithInsecure())
	if e != nil {
		panic(e)
	}
	noticeServiceClient = grpc2.NewNoticeServiceClient(conn)

}
func FilterInsert(notices []*grpc2.Notice) (*grpc2.FilterInsertReply, error) {
	background := context.Background()
	reply, e := noticeServiceClient.FilterInsert(background, &grpc2.FilterInsertRequest{Notices: notices})
	if e != nil {
		log.Logger.Error(e)
	}
	return reply, e
}
