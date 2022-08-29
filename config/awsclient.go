package config

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
)

func CreateConnection() (sess *session.Session) {
	log.Println("Tentando criar sessao na AWS...")
	return session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
}
