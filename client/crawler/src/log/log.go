package log

import (
	"go.uber.org/zap"
)

var production *zap.Logger
var Logger *zap.SugaredLogger

func init() {
	production, _ = zap.NewProduction()
	Logger = production.Sugar()
}
func Sync() {
	production.Sync()
}
