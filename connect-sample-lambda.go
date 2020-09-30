package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
	"time"
)

type Response struct {
	Answer string `json:"Answer" yaml:"Answer"`
}

func init() {
	masterLogger, err := zap.NewDevelopment()
	if err == nil {
		zap.ReplaceGlobals(masterLogger)
	}
}

//version variable which can be overidden at compile time
var (
	version = "dev-local-version"
	commit  = "none"
	date    = "unknown"
)

func HandleRequest(ctx context.Context, connectEvent events.ConnectEvent) (Response, error) {
	var logger = zap.S()

	logger.Infof("Version git Tag: %s Date: %s Commit: %s", version, date, commit)

	jLexEvent, _ := json.Marshal(connectEvent)
	logger.Info(string(jLexEvent))


	loc,_ := time.LoadLocation("Australia/Sydney")

	return Response{Answer : time.Now().In(loc).Format(time.Kitchen)}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
