package controllers

import (
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"

	"contact_api/interfaces"
	"contact_api/models"
)

type response struct {
	Status int              `json:"status"`
	Error  bool             `json:"error"`
	Data   interfaces.Model `json:"data"`
	Items  []models.Contact `json:"items"`
}

func Create(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, response{
			Status: 200,
			Error:  false,
			Data:   nil,
			Items:  nil,
		})
	}
}

func Index(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var contacts []models.Contact

		db.Find(&contacts)
		defer db.Close()

		res := response{
			Status: 200,
			Error:  false,
			Items:  contacts,
		}

		c.JSON(http.StatusOK, res)
	}
}

func GetContact(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

	}
}
