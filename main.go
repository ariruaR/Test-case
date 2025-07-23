package main

import (
	"time"

	gin "github.com/gin-gonic/gin"
)

type Subcription struct {
	ServiceName string    `json:"service_name"`
	Price       int       `json:"price"`
	UserID      string    `json:"user_id"`
	StartDate   time.Time `json:"start_date"`
}

//  GET /subscription
// POST /subcsription

func main() {
	g := gin.Default()

	g.GET("/test", func(ctx *gin.Context) {

	})
	g.Run(":8080")
}
