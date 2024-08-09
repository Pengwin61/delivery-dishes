package api

import (
	"delivery-dishes/internal/api/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func RunApi() {
	r := gin.Default()

	runHandlers(r)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func runHandlers(r *gin.Engine) {
	handlers.Ping(r)
}
