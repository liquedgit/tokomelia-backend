package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/liquedgit/tokoMeLia/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateNewUser is the resolver for the createNewUser field.
func (r *mutationResolver) CreateNewUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	if input.Password != input.ConfirmPassword {
		return nil, gqlerror.Errorf("Password not match")
	}

	newUser := &model.User{
		ID:       uuid.NewString(),
		Username: input.Username,
		Password: input.Password,
		Role:     input.Role,
	}

	res := r.DB.Save(newUser)

	return newUser, res.Error
}

// GetAllUser is the resolver for the GetAllUser field.
func (r *queryResolver) GetAllUser(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	return users, r.DB.Find(&users).Error
}

// GetUser is the resolver for the GetUser field.
func (r *queryResolver) GetUser(ctx context.Context, userid string) (*model.User, error) {
	var user *model.User
	return user, r.DB.First(&user, "username = ?", userid).Error
}

// LoginAccount is the resolver for the LoginAccount field.
func (r *queryResolver) LoginAccount(ctx context.Context) (*model.LoginResponse, error) {
	panic(fmt.Errorf("not implemented: LoginAccount - LoginAccount"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
