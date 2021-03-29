package dynamo

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func parseResult(result *dynamodb.ScanOutput) []Item {

	var rvItems []Item
	for _, i := range result.Items {
		item := Item{}

		err := dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			log.Fatalf("Got error unmarshalling: %s", err)
		}
		rvItems = append(rvItems, item)

	}
	return rvItems
}

func GetData(start, end string) []Item {
	// log.Println(data)
	region := os.Getenv("DYNAMO_REGION")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		log.Panicln(err)
	}
	svc := dynamodb.New(sess)
	tableName := os.Getenv("DYNAMO_TABLE")
	queryCol, projCol := "Date", "Data"
	startVal := expression.Value(start)
	endVal := expression.Value(end)
	exprGe := expression.Name(queryCol).GreaterThanEqual(startVal)
	exprLe := expression.Name(queryCol).LessThanEqual(endVal)
	filt := exprGe.And(exprLe)
	proj := expression.NamesList(expression.Name(projCol), expression.Name(queryCol))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		log.Fatalf("Got error building expression: %s", err)
	}
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableName),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
	}
	return parseResult(result)

}
