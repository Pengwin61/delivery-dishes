package handlers

import (
	"delivery-dishes/internal/db"
	"delivery-dishes/internal/logs"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Dish struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var (
	l *logs.Log
	d *db.ClientRoach
)

func Init(database *db.ClientRoach, logs *logs.Log) {
	l = logs
	d = database
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func AddDish(c *gin.Context) {
	var dish Dish
	err := c.BindJSON(&dish)
	if err != nil {
		l.Logger.DPanic("error bind json", zap.Error(err))
	}

	d.Con.Create(&Dish{Name: dish.Name, Price: dish.Price})

	c.JSON(http.StatusOK, gin.H{
		"message": "блюдо добавлено",
		"dish":    dish.Name,
	})
}

func GetDish(c *gin.Context) {
	var dish Dish
	err := c.BindJSON(&dish)
	if err != nil {
		l.Logger.DPanic("error bind json", zap.Error(err))
	}

	d.Con.Model(&Dish{}).Limit(10).Find(&Dish{})

	if result := d.Con.Where("name = ?", dish.Name).First(&dish); result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "нет такого блюда",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "получили блюдо",
			"dish":    dish.Name,
			"price":   dish.Price,
		})
	}
}
