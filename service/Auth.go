package service

import (
	"context"
	"github.com/liquedgit/tokoMeLia/graph/model"
	"github.com/liquedgit/tokoMeLia/helper"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

func UserLogin(ctx context.Context, username string, password string) (*model.LoginResponse, error) {
	getUser, err := UserGetByUsername(ctx, username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &gqlerror.Error{Message: "Invalid username or your account may not be activated"}
		}
		return nil, err
	}

	if err := helper.ComparePassword(getUser.Password, password); err != nil {
		return nil, err
	}

	token, err := GenerateToken(ctx, getUser.ID, getUser.Username, getUser.Role)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		Message: "success",
		Token:   token,
	}, nil

}
