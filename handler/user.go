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

func (h *userHandler) Login(c *gin.Context) {
	//user memasukan input (email & password)
	//input ditangkap handler
	//mapping dari input user ke input struct
	//input struct pasing service
	//di service mencari dengan bantuan repository user dengan email x
	//mencocokan password
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMassage := gin.H{"errors": errors}

		response := helper.APIResponse("Login filed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMassage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login filed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	formatter := user.FormatUser(loggedinUser, "tokentokentokentoken")

	response := helper.APIResponse("Successfuly loggedin", http.StatusOK, "succsess", formatter)
	c.JSON(http.StatusOK, response)
}
