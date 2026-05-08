package main

import (
	"net/http"

	"example.com/maphile/intializers"
	"github.com/gin-gonic/gin"
)


func init() {
	intializers.LoadEnvfile()
	intializers.Database_connection()
}

func main() {
	server := gin.Default()
    server.GET("/ping" , func(context *gin.Context){
		context.JSON(http.StatusOK , gin.H{"message" : "pong"})
	})
	server.Run()
}