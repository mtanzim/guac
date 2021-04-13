package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mtanzim/guac/server/services"
	"github.com/mtanzim/guac/server/utils"
)

type MyEvent struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func handler(ctx context.Context, event MyEvent) (string, error) {

	start := event.Start
	end := event.End

	if err := utils.ValidateQueryDate(start, end); err != nil {
		return "", err
	}

	rv := services.DataService(start, end)
	b, err := json.MarshalIndent(rv, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil

}

func main() {
	lambda.Start(handler)
}
