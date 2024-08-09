package api

import (
	"delivery-dishes/internal/api/handlers"
	"delivery-dishes/internal/db"
	"delivery-dishes/internal/logs"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const HTTP_PORT = ":80"

func RunApi() {

	d := db.NewClientCockroach()
	l := logs.InitLogs()

	fmt.Println("CURRNET_MIGRATION_DATABASE:", d.Con.Migrator().CurrentDatabase())

	d.Con.AutoMigrate(&handlers.Dish{})

	l.Logger.Info("start app mgs", zap.String("port", HTTP_PORT))

	handlers.Init(d, l)
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	runHandlers(r)

	err := r.Run(HTTP_PORT)
	if err != nil {
		log.Fatal(err)
	}
}

func runHandlers(r *gin.Engine) {

	v1 := r.Group("/v1")
	{
		v1.POST("/dish", handlers.AddDish)
		v1.GET("/dish", handlers.GetDish)
		v1.GET("/ping", handlers.Ping)

	}
}
