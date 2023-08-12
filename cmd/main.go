package main

import (
	"net/http"
	"redis-spin/internal"

	"github.com/gin-gonic/gin"
)


func main() {

	server := gin.Default()

	rdb := internal.OpenDB()

	server.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		result, err := internal.GetEntry(id, rdb)
		if err != nil {
			ctx.JSON(400, gin.H {"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, result)
	})

	server.POST("/add", func(ctx *gin.Context) {
		var reminder internal.Reminder
		ctx.ShouldBindJSON(&reminder)
		result, err := internal.AddToDB(rdb, &reminder)
		if err != nil {
			ctx.JSON(400, gin.H {"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, result)
	})

	server.DELETE("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		err := internal.DeleteEntry(id, rdb)
		if err != nil {
			ctx.JSON(400, gin.H {"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusNoContent, gin.H {})
	})


	server.Run(":8080")


}