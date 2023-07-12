package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/liquedgit/tokoMeLia/Database"
	"github.com/liquedgit/tokoMeLia/graph/model"
	"github.com/liquedgit/tokoMeLia/helper"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func UserCreate(ctx context.Context, input model.NewUser) (*model.User, error) {
	if input.Password != input.ConfirmPassword {
		return nil, &gqlerror.Error{Message: "Error password is not matched"}
	}
	db := Database.GetInstance()
	input.Password = helper.HashPasswords(input.Password)

	user := &model.User{
		ID:       uuid.NewString(),
		Username: input.Username,
		Password: input.Password,
		Role:     input.Role,
		Email:    input.Email,
	}
	res := db.Save(user)
	return user, res.Error
}

func UserGetByUsername(ctx context.Context, username string) (*model.User, error) {
	db := Database.GetInstance()
	var user *model.User
	return user, db.First(&user, "username = ?", username).Error
}

func UserGetByEmail(ctx context.Context, email string) (*model.User, error) {
	db := Database.GetInstance()
	var user *model.User
	return user, db.First(&user, "email = ?", email).Error
}
