package repository

import (
	"log"

	"br.com.eguadorodrigo/funcgoaws/config"
	"br.com.eguadorodrigo/funcgoaws/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
)

const TABLE_NAME = "toggle"

func GetItem(context *gin.Context) (response model.Toogle) {
	log.Printf("GetItem via DynamoDB invocado")

	ToggleName, OK := context.GetQuery("toggleName")

	if !OK {
		panic("Consulta não realizada, ToggleName não informado.")
	}

	log.Printf("ToggleName válido, consultando...")

	result, err := config.CreateConnectionDynamoDB().GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"toggleName": {
				S: aws.String(ToggleName),
			},
		},
	})

	log.Printf("Retornou resultado.")

	if err != nil {
		log.Fatalf("Erro ao chamar o método GetItem: %s", err)
	}

	if result.Item == nil {
		log.Fatalf("Não encontrado toggle com nome: %s", ToggleName)
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &response)

	if err != nil {
		log.Panic(err)
	}

	return response
}

func PostItem(context *gin.Context) string {

	marshalled, err := dynamodbattribute.MarshalMap(context.Request.Body)

	if err != nil {
		log.Fatalf("Erro ao converter objeto Toggle: %s", err)
	}

	toggle := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item:      marshalled,
	}

	_, err = config.CreateConnectionDynamoDB().PutItem(toggle)

	if err != nil {
		log.Fatalf("Erro ao cadastrar Toggle: %s", err)
	}

	return "Toggle cadastrado com sucesso."
}
