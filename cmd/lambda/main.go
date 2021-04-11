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
	Start string `json:"start"`
	End   string `json:"end"`
}

func HandleRequest(ctx context.Context, evt MyEvent) (string, error) {
	if err := utils.ValidateQueryDate(evt.Start, evt.End); err != nil {
		return fmt.Sprintf("Error: %s!", err.Error()), nil
	}
	rv := services.DataService(evt.Start, evt.End)
	b, err := json.MarshalIndent(rv, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error: %s!", err.Error()), nil
	}
	return string(b), nil

}

func main() {
	lambda.Start(HandleRequest)
}
