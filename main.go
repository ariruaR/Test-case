package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	models "exampler/components/models"

	_ "exampler/docs"

	gin "github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// @title Subscription API
// @version 1.0
// @description API для управления подписками и подсчёта их стоимости
// @host localhost:8080
// @BasePath /

// GetAllSubs godoc
// @Summary Получить все подписки
// @Tags subscriptions
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /subs [get]
func GetAllSubs(ctx *gin.Context) {
	subs, err := models.GetSubs(db)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		log.Println("Ошибка запроса")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"subscription": subs})
}

// CreateSub godoc
// @Summary Создать подписку
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param sub body models.Subcription true "Данные подписки"
// @Success 200 {object} models.Subcription
// @Failure 400 {object} map[string]interface{}
// @Router /subs [post]
func CreateSub(ctx *gin.Context) {
	var sub models.Subcription
	if err := ctx.ShouldBindJSON(&sub); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		log.Println("Ошибка запроса")
		return
	}
	err := models.CreateNewSub(sub.ServiceName, sub.Price, sub.UserID, sub.StartDate, db)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "Ошибка при добавлении информации в бд"})
		return
	}
	ctx.JSON(http.StatusOK, sub)
}

// GetSubsByUser godoc
// @Summary Получить подписки пользователя
// @Tags subscriptions
// @Produce json
// @Param userID path int true "ID пользователя"
// @Success 200 {array} models.Subcription
// @Failure 502 {object} map[string]interface{}
// @Router /subs/{userID} [get]
func GetSubsByUser(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.Param("userID"))
	subs, err := models.GetSubsById(db, userID)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, subs)
}

// UpdateSub godoc
// @Summary Изменить подписку пользователя
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param userID path int true "ID пользователя"
// @Param data body ChangeData true "Изменяемое поле"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /subs/{userID} [put]
func UpdateSub(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.Param("userID"))
	var data ChangeData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	res := models.UpdateSubs(db, userID, data.Field, data.Value)
	if res {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Ресурс успешно обновлен",
			"userID":  userID,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": data})
	}
}

// DeleteSub godoc
// @Summary Удалить подписку пользователя
// @Tags subscriptions
// @Produce json
// @Param userID path int true "ID пользователя"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /subs/{userID} [delete]
func DeleteSub(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.Param("userID"))
	err := models.DeleteSubs(db, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка удаления"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Ресурс успешно удален",
	})
}

// GetTotalPrice godoc
// @Summary Получить суммарную стоимость подписок
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param userID path string true "ID пользователя"
// @Param data body Data true "Фильтры: название сервиса, даты"
// @Success 200 {object} map[string]interface{}
// @Failure 502 {object} map[string]interface{}
// @Router /subs/price/{userID} [post]
func GetTotalPriceHandler(ctx *gin.Context) {
	userID := ctx.Param("userID")
	var data Data
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		log.Println("Ошибка запроса")
		return
	}
	log.Printf("POST /subs/price/%s: received data: %+v", userID, data)
	startDate, _ := time.Parse("2006-01-02", data.StartDate)
	endDate, _ := time.Parse("2006-01-02", data.EndDate)
	totalPrice, err := models.GetTotalPrice(db, userID, data.ServiceName, startDate, endDate)
	if err != nil {
		log.Printf("Error calculating total price: %v", err)
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err})
		return
	}
	log.Printf("Total price for userID=%s, serviceName=%s, startDate=%s, endDate=%s: %d", userID, data.ServiceName, startDate, endDate, totalPrice)
	ctx.JSON(http.StatusOK, gin.H{
		"Total price": totalPrice,
		"Start Date":  startDate,
		"End Date":    endDate,
	})
}

func main() {
	g := gin.Default()
	db = models.CreateDB()
	g.GET("/subs", GetAllSubs)
	g.POST("/subs", CreateSub)
	g.GET("/subs/:userID", GetSubsByUser)
	g.PUT("/subs/:userID", UpdateSub)
	g.DELETE("/subs/:userID", DeleteSub)
	g.POST("/subs/price/:userID", GetTotalPriceHandler)
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g.Run(":8080")
}

var db gorm.DB

type ChangeData struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}
type Data struct {
	ServiceName string `json:"service_name"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}
