package routes

import (
	"net/http"

	"github.com/FacelessFK/Event_booking/models"
	"github.com/FacelessFK/Event_booking/utils"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
	}

	err = user.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
	}

	err = user.ValidateCredentials()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
