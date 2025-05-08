package controller

import (
	"example/hello/database"
	"example/hello/models"
	"example/hello/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindUser(c *gin.Context) {

	var users []models.User

	database.DB.Find(&users)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "List data user",
		Data:    users,
	})
}
