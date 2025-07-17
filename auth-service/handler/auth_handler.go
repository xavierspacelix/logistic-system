package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/xavierspacelix/logistic-system/auth-service/model"
	"github.com/xavierspacelix/logistic-system/auth-service/repository"
	"github.com/xavierspacelix/logistic-system/auth-service/service"
	"github.com/xavierspacelix/logistic-system/auth-service/utils"
)

type RegisterRequest struct {
	MSISDN   string `json:"msisdn" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginRequest struct {
	MSISDN   string `json:"msisdn" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register godoc
// @Summary Register new user
// @Description Register with MSISDN and username
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.User true "Register Request"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /register [post]
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	if !strings.HasPrefix(req.MSISDN, "62") {
		utils.ResponseError(c, http.StatusBadRequest, "Invalid MSISDN format", gin.H{
			"msisdn": "must start with 62",
		})
		return
	}

	if _, err := repository.FindUserByUsernameOrMSISDN(req.Username, req.MSISDN); err == nil {
		utils.ResponseError(c, http.StatusBadRequest, "Username or MSISDN already exists", nil)
		return
	}

	hashed, _ := service.HashPassword(req.Password)

	user := &model.User{
		ID:       uuid.New(),
		Name:     req.Name,
		MSISDN:   req.MSISDN,
		Username: req.Username,
		Password: hashed,
	}

	if err := repository.CreateUser(user); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Failed to create user", err.Error())
		return
	}

	utils.ResponseCreated(c, "User registered successfully", gin.H{
		"id":       user.ID,
		"username": user.Username,
		"msisdn":   user.MSISDN,
		"name":     user.Name,
	})
}

// Login godoc
// @Summary Login user
// @Description Authenticate user with MSISDN and password
// @Tags auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "Login Request"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /login [post]
func Login(c *gin.Context) {
	var req struct {
		MSISDN   string `json:"msisdn" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	user, err := repository.FindUserByMSISDN(req.MSISDN)
	if err != nil || !service.CheckPasswordHash(req.Password, user.Password) {
		utils.ResponseError(c, http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}

	token, _ := service.GenerateJWT(user.ID.String())
	utils.ResponseOK(c, "Login successful", gin.H{
		"token": token,
	})
}

// Me godoc
// @Summary Get current user profile
// @Description Get information about the logged in user
// @Tags auth
// @Security BearerAuth
// @Produce json
// @Success 200 {object} utils.SuccessResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /me [get]
func Me(c *gin.Context) {
	userID := c.GetString("user_id")

	user, err := repository.FindUserByID(userID)
	if err != nil {
		utils.ResponseError(c, http.StatusNotFound, "User not found", nil)
		return
	}

	utils.ResponseOK(c, "User profile fetched", gin.H{
		"id":       user.ID,
		"username": user.Username,
		"name":     user.Name,
		"msisdn":   user.MSISDN,
	})
}
