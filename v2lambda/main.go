package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type Payload struct {
	To  string `json:"to"`
	Msg string `json:"msg"`
}

func HandleRequest(ctx context.Context, event *Payload) (*string, error) {
	fmt.Println(event)
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}
	resp, err := SendSMS(*event)
	if err != nil {
		return nil, err
	}
	message := fmt.Sprintf("Hello %s!", string(resp))
	return &message, nil
}

func main() {
	lambda.Start(HandleRequest)
}

func SendSMS(payload Payload) (resp []byte, err error) {

	from := os.Getenv("from")
	username := os.Getenv("username")
	password := os.Getenv("password")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: username,
		Password: password,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(payload.To)
	params.SetFrom(from)
	params.SetBody(payload.Msg)

	t, err := client.Api.CreateMessage(params)
	if err != nil {
		return nil, err
	}
	return json.Marshal(*t)
}
