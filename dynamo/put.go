package dynamo

import (
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const category = "coding"

type Item struct {
	Category string // primary key
	Date     string // sort key
	Data     interface{}
}

func putData(item *Item, svc *dynamodb.DynamoDB, tableName string, wg *sync.WaitGroup) {
	defer wg.Done()
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Println("Got error marshalling new waka item:")
		log.Fatalln(err.Error())
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	_, err = svc.PutItem(input)
	if err != nil {
		log.Println("Got error calling PutItem:")
		log.Fatalln(err.Error())
	}
	log.Println("Successfully added/updated date " + item.Date + " to table " + tableName)
}

func PutData(data map[string]interface{}) {
	// log.Println(data)
	region := os.Getenv("DYNAMO_REGION")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		log.Fatalln(err)
	}
	svc := dynamodb.New(sess)
	tableName := os.Getenv("DYNAMO_TABLE")
	var wg sync.WaitGroup
	// create the input configuration instance
	for k, v := range data {
		item := Item{category, k, v}
		wg.Add(1)
		go putData(&item, svc, tableName, &wg)
	}
	wg.Wait()

}
