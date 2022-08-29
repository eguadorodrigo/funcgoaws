package router

import (
	"log"

	"br.com.eguadorodrigo/funcgoaws/controller"
	"github.com/gin-gonic/gin"
)

func ToggleRouter(route *gin.Engine){
	log.Printf("Aplicando regra de roteamento")
	route.GET("/toggle", controller.Get)
	route.POST("/toggle", controller.Post)
}