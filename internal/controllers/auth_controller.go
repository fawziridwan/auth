package controllers

import (
	"net/http"
	"time"

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

	responses.SuccessResponse(
		ctx,
		http.StatusCreated,
		"User registered successfully",
		models.RegisterResponse{
			StatusCode: http.StatusCreated,
			Status:     true,
			Message:    "User registered successfully",
			User: models.RegisterData{
				Name:      req.Name,
				Email:     req.Email,
				Password:  req.Password, // Note: Password should not be returned in production
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	)
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, user, err := c.authService.Login(&req)
	if err != nil {
		responses.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	userData := models.UserData{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password, // Note: Password should not be returned in production
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	responses.SuccessResponse(ctx, http.StatusOK, "Login successful", models.AuthResponse{
		Token: token,
		User:  userData,
	})
}
