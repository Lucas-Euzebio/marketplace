package controller

import (
	"ecommerce/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Getuser(context *gin.Context) {
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user})
}
