package main

import (
	"redis-spin/internal"

	"github.com/gin-gonic/gin"
)


func main() {

	server := gin.Default()

	rdb := internal.OpenDB()

	server.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		result, err := rdb.Get(id).Result()
		if err != nil {
			ctx.JSON(400, gin.H {"error": err.Error()})
			return
		}
		ctx.JSON(200, result)
	})

	server.POST("/add", func(ctx *gin.Context) {

	})

	server.DELETE("/:id", func(ctx *gin.Context) {

	})


	server.Run(":8080")


}