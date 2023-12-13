package talentcontroller

import (
	"golang_talent/helper"
	"golang_talent/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var talent []models.Talent
	models.DB.Find(&talent)
	helper.ResponseWithData(c, http.StatusOK, "0000", talent)
}

func Create(c *gin.Context) {

}
func Update(c *gin.Context) {

}
func Delete(c *gin.Context) {

}
func Show(c *gin.Context) {
	var talent models.Talent
	// Extract ID from the request parameter
	id := c.Param("id")

	// Attempt to find the talent record in the database
	if err := models.DB.First(&talent, id).Error; err != nil {
		switch err {
		// If the talent record is not found
		case gorm.ErrRecordNotFound:
			helper.ResponseWithError(c, http.StatusNotFound, "404", "data not found")
			return
		default:
			helper.ResponseWithError(c, http.StatusInternalServerError, "500", err.Error())
			return
		}
	}
	// Respond with the fetched talent data
	helper.ResponseWithData(c, http.StatusOK, "0000", talent)

}
