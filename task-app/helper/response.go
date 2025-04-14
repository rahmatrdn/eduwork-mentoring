package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(c *gin.Context, data interface{}, message string) {
    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": message,
        "data":    data,
    })
}

func ResponseCreated(c *gin.Context, data interface{}, message string) {
    c.JSON(http.StatusCreated, gin.H{
        "status":  "success",
        "message": message,
        "data":    data,
    })
}

func ResponseError(c *gin.Context, code int, message string) {
    c.JSON(code, gin.H{
        "status":  "error",
        "message": message,
    })
}
