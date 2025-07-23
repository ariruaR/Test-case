package main

import (
	"net/http"

	models "exampler/components/models"

	gin "github.com/gin-gonic/gin"
)

//  GET /subs
// GET /subs/{userID}
// POST /subs
// PUT /subs/{userID}
// DELETE /subs/{userID}

func main() {
	g := gin.Default()

	g.GET("/subs", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"subscription": "ok"})
	})
	g.POST("/subs", func(ctx *gin.Context) {
		var sub models.Subcription
		if err := ctx.ShouldBindJSON(&sub); err != nil {
			ctx.Status(http.StatusBadRequest)
		}
		ctx.JSON(http.StatusOK, sub)
	})

	g.GET("/subs/:userID", func(ctx *gin.Context) {
		userID := ctx.Param("userID")
		//! получаем данные из БД
		ctx.JSON(http.StatusOK, gin.H{"userID": userID})
	})

	g.PUT("/subs/:userID", func(ctx *gin.Context) {
		userID := ctx.Param("userID")
		// ! логика обновления ресурса по ID
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Ресурс успешно обновлен",
			"userID":  userID,
		})
	})
	g.DELETE("/subs/:userID", func(ctx *gin.Context) {
		userID := ctx.Param("userID")
		// ! логика обновления ресурса по ID
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Ресурс успешно удален",
			"userID":  userID,
		})
	})

	g.Run(":8080")
}
