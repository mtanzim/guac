package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mtanzim/guac/dynamo"
	"github.com/mtanzim/guac/processData"
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
	data := dynamo.GetData(evt.Start, evt.End)
	dailyStats := processData.DailyTotal(data)
	_, _ = processData.GetDateRange(dailyStats)

	_ = processData.LanguageSummary(data)

	return fmt.Sprintf("Hello %s!", data), nil
}

func main() {
	lambda.Start(HandleRequest)
}
