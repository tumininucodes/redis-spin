package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)


func main() {

	server := gin.Default()

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	ping, err := client.Ping(&gin.Context{}).Result()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("redis client >> ", ping)


	server.GET("/:id", func(ctx *gin.Context) {

	})

	server.POST("/add", func(ctx *gin.Context) {

	})

	server.DELETE("/:id", func(ctx *gin.Context) {

	})


	server.Run(":8080")
}