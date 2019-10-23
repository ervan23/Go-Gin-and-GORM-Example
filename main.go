package main

import (
	"contact_api/controllers"
	"contact_api/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := gin.Default()

	db, err := gorm.Open("mysql", "root:@(localhost)/contact_api?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	if !db.HasTable(&models.Contact{}) {
		db.AutoMigrate(&models.Contact{})
	}

	r.GET("/", controllers.Index)
	r.GET("/:Id", controllers.GetContact)
	r.POST("/", controllers.Create)

	r.Run(":2309")
}