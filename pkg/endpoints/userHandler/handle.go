package userHandler

import (
	"main/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserValidator UserValidator
	// service UserServices
}

func NewHandler(userValidator UserValidator) Handler {

	return Handler{UserValidator: userValidator}
}

func (h Handler) AddUser(c *gin.Context) {

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.UserValidator.validator.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// h.service.AddUserService

	c.JSON(http.StatusOK, gin.H{"message": "User checked"})

}
