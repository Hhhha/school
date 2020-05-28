package mongo

import (
	"context"
	"crawler/client/crawler/log"
	grpc "crawler/grpc/server"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const addr = "mongodb://47.97.90.202:27017"

var Collection *mongo.Collection
var Ctx context.Context

func init() {
	client, e := mongo.NewClient(options.Client().ApplyURI(addr))
	if e != nil {
		log.Logger.Error(e)
	}
	Ctx = context.Background()
	e = client.Connect(Ctx)
	if e != nil {
		log.Logger.Error(e)
	}
	Collection = client.Database("crawler").Collection("notice")
}
func Insert(notices []*grpc.Notice) (*mongo.InsertManyResult, error) {
	inserts := make([]interface{}, 0)
	for _, value := range notices {
		inserts = append(inserts, value)
	}
	result, e := Collection.InsertMany(Ctx, inserts)
	return result, e
}
func FindOneByTitle(title string) (grpc.Notice, error) {
	notice := grpc.Notice{}
	err := Collection.FindOne(Ctx, bson.M{
		"title": primitive.Regex{Pattern: title},
	}).Decode(&notice)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return notice, err
		}
		panic(err)
	}
	log.Logger.Info("query:", notice)
	return notice, nil
}
