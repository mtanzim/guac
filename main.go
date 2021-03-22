package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mtanzim/guac/wakaApi"
)

type Item struct {
	Date string
	Data interface{}
}

func main() {
	data := wakaApi.TransformData()
	// log.Println(data)
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		log.Panicln(err)
	}
	svc := dynamodb.New(sess)
	tableName := os.Getenv("DYNAMO_TABLE")
	// create the input configuration instance
	for k, v := range data {
		item := Item{k, v}
		av, err := dynamodbattribute.MarshalMap(item)
		if err != nil {
			log.Println("Got error marshalling new waka item:")
			log.Panicln(err.Error())
		}
		// log.Println(av)
		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String(tableName),
		}
		_, err = svc.PutItem(input)
		if err != nil {
			log.Println("Got error calling PutItem:")
			log.Panicln(err.Error())
		}
		log.Println("Successfully added date " + item.Date + " to table " + tableName)
	}

}
