package services

import (
	"errors"
	"time"

	"github.com/beevik/guid"
	"github.com/dgrijalva/jwt-go"
	"test.com/apiTest/configurations"
	"test.com/apiTest/models"
	"test.com/apiTest/repositories"
)

type UserSignedInDto struct {
	Id       string
	UserName string
	Email    string
	JwtToken string
}

type SignInService interface {
	SignIn(username string, password string) (*UserSignedInDto, error)
}

type SignInSeviceImplementation struct {
	UsersRepository repositories.UsersRepository
	Config          configurations.Configuration
}

func (service *SignInSeviceImplementation) SignIn(username string, password string) (*UserSignedInDto, error) {
	user, err := service.UsersRepository.GetByUserName(username)

	if !user.ComparePassword(password) {
		return nil, errors.New("invalid credentials")
	}

	token, err := createJwtToken(user, service.Config.Secret)
	dto := &UserSignedInDto{
		Email:    user.Email,
		Id:       user.Id,
		JwtToken: token,
		UserName: user.UserName,
	}

	return dto, err
}

func createJwtToken(user models.User, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{
		Id:        guid.NewString(),
		Subject:   user.Id,
		ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}
