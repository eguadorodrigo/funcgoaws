package config

import (
	"log"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateConnectionDynamoDB()(svc *dynamodb.DynamoDB){
	log.Printf("Tentando conexão com o DynamoDB.")
	return dynamodb.New(CreateConnection())
}