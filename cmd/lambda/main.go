package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mtanzim/guac/server/services"
	"github.com/mtanzim/guac/server/utils"
)

type MyEvent struct {
	QueryStringParameters struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"queryStringParameters"`
}

func HandleRequest(ctx context.Context, evt MyEvent) (string, error) {

	start := evt.QueryStringParameters.Start
	end := evt.QueryStringParameters.End
	if err := utils.ValidateQueryDate(start, end); err != nil {
		return fmt.Sprintf("Error: %s!", err.Error()), nil
	}
	rv := services.DataService(start, end)
	b, err := json.MarshalIndent(rv, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error: %s!", err.Error()), nil
	}
	return string(b), nil

}

func main() {
	lambda.Start(HandleRequest)
}
