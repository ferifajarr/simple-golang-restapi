package helper

import (
	"github.com/gin-gonic/gin"
)

var ResponseMap = map[string]string{
	"0000": "Success",
	"500":  "Internal Server Error",
	"400":  "Bad Request",
	"403":  "Unauthorized",
	"404":  "Not Found",
}

func ResponseWithError(c *gin.Context, code int, responseCode string, errMsg string) {
	c.AbortWithStatusJSON(code, gin.H{
		"responseCode":        responseCode,
		"responseDescription": ResponseMap[responseCode],
		"responseMessage":     errMsg,
	})
}
func ResponseWithData(c *gin.Context, code int, responseCode string, data interface{}) {
	c.JSON(code, gin.H{
		"responseCode":        responseCode,
		"responseDescription": ResponseMap[responseCode],
		"data":                data,
	})
}
