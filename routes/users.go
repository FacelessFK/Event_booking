package routes

import (
	"net/http"

	"github.com/FacelessFK/Event_booking/models"
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
