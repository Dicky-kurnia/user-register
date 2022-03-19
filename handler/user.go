package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
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

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	// ada input email dari user
	// input email di mapping ke struct input
	// struct input di passing ke service
	// service akan manggil repository - emqil sudah ada atau belum
	// repository - db

	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMassage := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMassage := gin.H{"errors": "Server error"}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "error", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {

	//input dari user
	//simpan gambarnya di folder "images/"
	//di service kita ambil repo
	//JWT (sementara hardcode, seakan akan user yang login ID = 1)
	//repo ambil data user yang ID = 1
	//repo update data user simpan lokasi file

	file, err := c.FormFile("avtar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avtar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := 1

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avtar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avtar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar successfully uploaded", http.StatusOK, "success`", data)

	c.JSON(http.StatusOK, response)
}
