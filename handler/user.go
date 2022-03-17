package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUSer(c *gin.Context) {
	// tangkap input dari user
	// map input dari uiser ke struct RegisterUserInput
	//struct di atas kita pasing sebagai parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMassage := gin.H{"errors": errors}
		response := helper.APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "succsess", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register Account  Failed", http.StatusBadRequest, "succsess", nil)
		c.JSON(http.StatusOK, response)
	}

	// token, err := h.jwtService.GenerateToken()

	formatter := user.FormatUser(newUser, "tokentokentokentoken")

	response := helper.APIResponse("Accound has been registered", http.StatusOK, "succsess", formatter)
	c.JSON(http.StatusOK, response)
}
