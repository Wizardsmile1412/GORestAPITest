package routes

import (
	"net/http"

	"example.com/restAPIProject/models"
	"example.com/restAPIProject/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"messege": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messege": "Could not save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"messege": "User created successfully"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindBodyWithJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"messege": "Could not parse request data."})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"messege": "Could not authenticate user."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messege": "Could not authenticate user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
