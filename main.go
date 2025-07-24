package main

import (
	"log"
	"net/http"
	"strconv"

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

	db := models.CreateDB()

	g.GET("/subs", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"subscription": "ok"})
	})
	g.POST("/subs", func(ctx *gin.Context) {
		var sub models.Subcription
		if err := ctx.ShouldBindJSON(&sub); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println("Ошибка запроса")
		}
		err := models.CreateNewSub(sub.ServiceName, sub.Price, sub.UserID, sub.StartDate, db)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadGateway, gin.H{"message": "Ошибка при добавлении информации в бд"})
		}
		ctx.JSON(http.StatusOK, sub)
	})

	g.GET("/subs/:userID", func(ctx *gin.Context) {
		userID, _ := strconv.Atoi(ctx.Param("userID"))

		subs, err := models.GetSubsById(db, userID)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadGateway, gin.H{"error": err})
		}
		ctx.JSON(http.StatusOK, subs)
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
