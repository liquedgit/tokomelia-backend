package service

import (
	"context"
	"encoding/base64"
	"errors"
	"github.com/google/uuid"
	"github.com/liquedgit/tokoMeLia/Database"
	"github.com/liquedgit/tokoMeLia/graph/model"
	"github.com/liquedgit/tokoMeLia/helper"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"net/url"
	"time"
)

func VerifyEmail(ctx context.Context, encodedToken string) (*model.DefaultResponse, error) {
	unescapedToken, err := url.QueryUnescape(encodedToken)
	tokenByte, err := base64.StdEncoding.DecodeString(unescapedToken)
	token := string(tokenByte)
	if token == "" {
		return nil, errors.New("Error While Verifying Email")
	}

	validate, err := JwtValidate(ctx, token)
	if err != nil || !validate.Valid {
		return nil, errors.New("Error while verifying email")
	}

	customClaim, _ := validate.Claims.(*JwtCustomClaim)
	username := customClaim.Username
	user, err := UserGetByUsername(ctx, username)
	if err != nil {
		panic(customClaim.Username)
		return nil, errors.New(err.Error())
	}
	user.IsVerified = true
	db := Database.GetInstance()
	db.Save(&user)
	return &model.DefaultResponse{Message: "Succesfully verified account"}, db.Save(&user).Error
}

func UserCreate(ctx context.Context, input model.NewUser) (*model.User, error) {
	if input.Password != input.ConfirmPassword {
		return nil, &gqlerror.Error{Message: "Error password is not matched"}
	}
	db := Database.GetInstance()
	input.Password = helper.HashPasswords(input.Password)

	user := &model.User{
		ID:         uuid.NewString(),
		Username:   input.Username,
		Password:   input.Password,
		Role:       input.Role,
		Email:      input.Email,
		CreatedAt:  time.Time{},
		IsVerified: false,
	}
	res := db.Save(&user)

	token, err := GenerateToken(ctx, user.ID, user.Username, user.Role)
	if err != nil {
		panic("Error Occured")
	}
	encodedToken := base64.StdEncoding.EncodeToString([]byte(token))
	escapedToken := url.QueryEscape(encodedToken)
	verifyEmailLink := "http://localhost:5173/verify/" + escapedToken
	SendMail(verifyEmailLink, input.Email)

	return user, res.Error
}

func UserGetByUsername(ctx context.Context, username string) (*model.User, error) {
	db := Database.GetInstance()
	var user *model.User
	return user, db.First(&user, "username = ? AND is_verified = true", username).Error
}

func UserGetByEmail(ctx context.Context, email string) (*model.User, error) {
	db := Database.GetInstance()
	var user *model.User
	return user, db.First(&user, "email = ? AND is_verified = true", email).Error
}

func UserGetByID(ctx context.Context, id string) (*model.User, error) {
	db := Database.GetInstance()
	var user *model.User
	return user, db.First(&user, "id = ?", id).Error
}
