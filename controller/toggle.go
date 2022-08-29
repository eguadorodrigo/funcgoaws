package controller

import (
	"log"
	"net/http"

	"br.com.eguadorodrigo/funcgoaws/repository"
	"github.com/gin-gonic/gin"
)

func Get(context *gin.Context) {
	log.Printf("Métogo GET invocado na Controller")
	context.IndentedJSON(http.StatusOK, repository.GetItem(context))
}

func Post(context *gin.Context) {
	log.Printf("Métogo POST invocado na Controller")
	context.IndentedJSON(http.StatusOK, repository.PostItem(context))
}
