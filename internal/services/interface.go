package services

import (
	"project-skbackend/internal/controllers/requests"
	"project-skbackend/internal/controllers/responses"
	"project-skbackend/internal/models"
	"project-skbackend/packages/utils"

	"github.com/google/uuid"
)

type (
	IUserService interface {
		CreateUser(req requests.CreateUserRequest) (*responses.UserResponse, error)
		GetUser(id uuid.UUID) (*responses.UserResponse, error)
		GetUsers(paginationReq models.Pagination) (*models.Pagination, error)
		DeleteUser(id uuid.UUID) error
	}

	IAuthService interface {
		Login(req requests.LoginRequest) (*responses.UserResponse, *utils.TokenHeader, error)
		Register(req requests.RegisterRequest) (*responses.UserResponse, *utils.TokenHeader, error)
		ForgotPassword(req requests.ForgotPasswordRequest) error
		ResetPassword(req requests.ResetPasswordRequest) error
		SendVerificationEmail(id uuid.UUID, token int) error
		VerifyToken(req requests.VerifyTokenRequest) error
		SendResetPasswordEmail(id uuid.UUID, token int) error
		RefreshAuthToken(token string) (*responses.UserResponse, *utils.TokenHeader, error)
	}

	IMailService interface {
		SendVerificationEmail(req requests.SendEmailRequest) error
	}
)
