package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nguyendangcuong201004/event-booking-api/models"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not signup user. Try again later"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User signup successfuly!", "user": user})
}
