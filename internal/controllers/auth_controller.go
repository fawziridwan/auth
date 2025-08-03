package controllers

import (
	"net/http"

	"github.com/fawziridwan/auth_module/internal/models"
	"github.com/fawziridwan/auth_module/internal/services"
	"github.com/fawziridwan/auth_module/internal/utils/responses"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req models.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.authService.Register(&req); err != nil {
		responses.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	responses.SuccessResponse(ctx, http.StatusCreated, "User registered successfully", nil)
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := c.authService.Login(&req)
	if err != nil {
		responses.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	responses.SuccessResponse(ctx, http.StatusOK, "Login successful", models.AuthResponse{Token: token})
}
