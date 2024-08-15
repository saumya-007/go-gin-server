package main

import (
	"time"

	"github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
	controller "github.com/saumya-007/go-gin-server/controller"
	ConnectMongo "github.com/saumya-007/go-gin-server/db-access"
)

func main() {
	ConnectMongo.ConnectMongo()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders: []string{
			"x-tenant-id",
			"x-workspace-id",
			"x-env-id",
			"accesstoken",
			"connection",
			"x-workspace-hash",
			"Content-Type",
		},
		MaxAge: 12 * time.Hour,
	}))
	solvedQuestionsRoutes(router)
	router.Run("localhost:7081")
}

func solvedQuestionsRoutes(router *gin.Engine) {
	router.GET("/solved-questions", controller.GetSolvedQuestions)
	router.POST("/solved-question", controller.AddSolvedQuestion)
	router.PUT("/solved-question/:id", controller.UpdateSolvedQuestion)
	router.DELETE("/solved-question/:id", controller.DeleteSolvedQuestion)
}
