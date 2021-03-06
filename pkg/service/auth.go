package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/alex-dev-master/chat-golang/pkg/model"
	"github.com/alex-dev-master/chat-golang/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"math/rand"
	"time"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (int64, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GetUser(username, password string) (model.User, error) {
	return s.repo.GetUser(username, generatePasswordHash(password))
}

func (s *AuthService) GenerateToken(user model.User) (string, error) {
	return getNewAccessToken(user.Id)
}

func (s *AuthService) CreateRefreshToken(user model.User) (string, error) {
	newRefreshToken, err := getNewRefreshToken()
	if err != nil {
		return "", err
	}

	err = s.repo.UpdateRefreshToken(user, newRefreshToken)
	if err != nil {
		return "", err
	}

	return newRefreshToken, err
}

func (s *AuthService) RefreshAccessToken(refreshToken string) (string, error) {

	user, err := s.repo.GetByRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}

	now := time.Now().Unix()
	if now > user.ExpiredToken {
		return "", errors.New("refresh token expired")
	}

	return getNewAccessToken(user.Id)
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func getNewAccessToken(userId int) (string, error) {
	accessTokenTTLStr := viper.GetString("auth.accessTokenTTL")
	accessTokenTTL, _ := time.ParseDuration(accessTokenTTLStr)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})

	return token.SignedString([]byte(signingKey))
}

func getNewRefreshToken() (string, error) {
	b := make([]byte, 32)

	a := rand.NewSource(time.Now().Unix())
	r := rand.New(a)

	_, err := r.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
