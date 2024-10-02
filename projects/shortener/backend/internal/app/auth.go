package app

import (
	"backend/model"
	"backend/pkg"
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var PasswordHashStrength = 10

func (a *app) hashPassword(password string) (string, error) {
	if err := pkg.Validator.ValidateVar(password, "required"); err != nil {
		return "", err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), PasswordHashStrength)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (a *app) comparePassword(hash, password string) bool {
	if err := pkg.Validator.ValidateMap(
		pkg.ValidationMapData{
			"hash":     hash,
			"password": password,
		}, pkg.ValidationMapRule{
			"hash":     "required",
			"password": "required",
		}); err != nil {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a *app) createJWTToken(userID int, email string) (string, error) {
	if err := pkg.Validator.ValidateMap(
		pkg.ValidationMapData{
			"userID": userID,
			"email":  email,
		}, pkg.ValidationMapRule{
			"userID": "required",
			"email":  "required",
		}); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userID,
			"email":   email,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	tokenString, err := token.SignedString([]byte(os.Getenv("SHORTENER_JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *app) ClaimJWTToken(tokenString string) (*model.JWTClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SHORTENER_JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid jwt token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("error parsing claims")
	}

	response := &model.JWTClaims{
		UserID: claims["user_id"].(float64),
		Email:  claims["email"].(string),
	}

	return response, nil
}

func (a *app) Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error) {
	if err := pkg.Validator.ValidateStruct(request); err != nil {
		return nil, err
	}

	user, err := a.store.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	if !a.comparePassword(user.Password, request.Password) {
		return nil, errors.New("invalid username or password")
	}

	jwtToken, err := a.createJWTToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	response := &model.LoginResponse{
		JWTToken: jwtToken,
	}

	return response, nil
}

func (a *app) Register(ctx context.Context, request *model.CreateUserRequest) error {
	if err := pkg.Validator.ValidateStruct(request); err != nil {
		return err
	}

	user, err := a.store.GetUserByEmail(ctx, request.Email)
	if err != nil && err != model.ErrDataNotFound {
		return err
	}
	if user != nil {
		return errors.New("user already exists")
	}

	hashPassword, err := a.hashPassword(request.Password)
	if err != nil {
		return err
	}

	request.Password = hashPassword
	return a.store.CreateUser(ctx, request)
}
