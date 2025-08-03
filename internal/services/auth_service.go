package services

import (
	"errors"
	"time"

	"github.com/fawziridwan/auth_module/internal/models"
	"github.com/fawziridwan/auth_module/internal/repositories"
	"github.com/fawziridwan/auth_module/internal/utils/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(user *models.RegisterRequest) error
	Login(credentials *models.LoginRequest) (string, error)
}

type authService struct {
	userRepo  repositories.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo repositories.UserRepository, jwtSecret string) AuthService {
	return &authService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *authService) Register(userReq *models.RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: string(hashedPassword),
	}

	return s.userRepo.Create(user)
}

func (s *authService) Login(credentials *models.LoginRequest) (string, error) {
	user, err := s.userRepo.FindByEmail(credentials.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := jwt.GenerateToken(user.ID, s.jwtSecret, 24*time.Hour)
	if err != nil {
		return "", err
	}

	return token, nil
}
