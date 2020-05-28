package main

import (
	"crawler/client/crawler/log"
	server "crawler/grpc/server"
	"crawler/service/notice/config"
	notice "crawler/service/notice/service"
	"google.golang.org/grpc"
	"net"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	listener, err := net.Listen("tcp", ":"+config.PORT)
	if err != nil {
		log.Logger.Error(err)
	}
	noticeServer := grpc.NewServer()
	server.RegisterNoticeServiceServer(noticeServer, &notice.NoticeServer{})
	err = noticeServer.Serve(listener)
	if err != nil {
		log.Logger.Error(err)
	}
}
