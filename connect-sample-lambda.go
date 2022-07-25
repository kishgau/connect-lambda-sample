package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

type MyEvent struct {
	Name string `json:"name"`
}

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
	loc, _ := time.LoadLocation("Australia/Sydney")

	userSelection := connectEvent.Details.ContactData.Attributes["UserSelection"]
	// return Response{Answer: time.Now().In(loc).Format(time.Kitchen)}, nil

	// Details.ContactData.Attributes.UserSelection)
	var answerText string
	switch userSelection {
	case "1":
		answerText = "Today's date is" + time.Now().In(loc).Format("01 January 2006")
	case "2":
		answerText = "The time now is" + time.Now().In(loc).Format(time.Kitchen)
	case "3":
		answerText = "Today is " + time.Now().In(loc).Weekday().String()

	}

	return Response{Answer: answerText}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
