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

func parseResult(result *dynamodb.QueryOutput) []Item {

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
	if region == "" {
		log.Fatalln("Please provide region")
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		log.Panicln(err)
	}
	svc := dynamodb.New(sess)
	tableName := os.Getenv("DYNAMO_TABLE")
	if region == "" {
		log.Fatalln("Please provide table name")
	}
	sortCol, projCol, keyCol := "Date", "Data", "Category"
	startVal := expression.Value(start)
	endVal := expression.Value(end)
	keyConditionEq := expression.KeyEqual(expression.Key(keyCol), expression.Value("coding"))
	sortCondition := expression.KeyBetween(expression.Key(sortCol), startVal, endVal)
	overallQuery := expression.KeyAnd(keyConditionEq, sortCondition)

	proj := expression.NamesList(expression.Name(projCol), expression.Name(sortCol))
	expr, err := expression.NewBuilder().WithKeyCondition(overallQuery).WithProjection(proj).Build()
	if err != nil {
		log.Fatalf("Got error building expression: %s", err)
	}

	params := &dynamodb.QueryInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		// FilterExpression:          expr.Filter(),
		ProjectionExpression:   expr.Projection(),
		KeyConditionExpression: expr.KeyCondition(),
		TableName:              aws.String(tableName),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Query(params)
	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
	}
	return parseResult(result)

}
