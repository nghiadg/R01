package auth

import (
	"errors"
	"r01/internal/repository"
	"r01/internal/utils/jwt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase struct {
	userRepo repository.IUserRepository
}

func NewLoginUsecase(userRepo repository.IUserRepository) *LoginUsecase {
	return &LoginUsecase{userRepo: userRepo}
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResult struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// errors
var (
	ErrInvalidPassword = errors.New("invalid password")
)

func (lu LoginUsecase) Execute(params LoginParams) (*LoginResult, error) {
	user, err := lu.userRepo.FindByEmail(params.Email)
	if err != nil {
		return nil, err
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))

	if err != nil {
		return nil, ErrInvalidPassword
	}

	// create token and return
	payload := map[string]interface{}{
		"email": user.Email,
		"id":    user.ID,
	}

	accessToken, err := jwt.GenerateToken(payload, time.Now().Add(time.Minute*24).Unix())

	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.GenerateToken(payload, time.Now().Add(time.Hour*24).Unix())

	if err != nil {
		return nil, err
	}

	return &LoginResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}
