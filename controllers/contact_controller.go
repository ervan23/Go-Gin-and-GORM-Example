package controllers

import (
	"net/http"

	"github.com/google/uuid"

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
		var contact models.Contact
		contact.Id = uuid.New().String()
		c.BindJSON(&contact)
		db.Create(&contact)

		c.JSON(http.StatusOK, response{
			Status: 200,
			Error:  false,
			Data:   contact,
			Items:  nil,
		})
	}
}

func Index(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var contacts []models.Contact

		db.Find(&contacts)

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
		var contact models.Contact
		db.Where("id = ?", c.Param("Id")).First(&contact)

		c.JSON(http.StatusOK, response{
			Status: 200,
			Error:  false,
			Data:   contact,
			Items:  nil,
		})
	}
}

func UpdateContact(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var contact, params models.Contact
		c.BindJSON(&params)
		db.Where("id = ?", c.Param("Id")).First(&contact)
		db.Model(&contact).Updates(models.Contact{
			Name:  params.Name,
			Phone: params.Phone,
			Email: params.Email,
		})

		c.JSON(http.StatusOK, response{
			Status: 200,
			Error:  false,
			Data:   contact,
			Items:  nil,
		})
	}
}

func DeleteContact(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var contact models.Contact
		db.Where("id = ?", c.Param("Id")).First(&contact)
		db.Delete(&contact)

		c.JSON(http.StatusOK, response{
			Status: 200,
			Error:  false,
			Data:   nil,
			Items:  nil,
		})
	}
}
