package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/liquedgit/tokoMeLia/middlewares"
	"time"

	"github.com/liquedgit/tokoMeLia/graph/model"
	"github.com/liquedgit/tokoMeLia/service"
)

// Chat is the resolver for the chat field.
func (r *chatDetailsResolver) Chat(ctx context.Context, obj *model.ChatDetails) (*model.ChatHeader, error) {
	return service.GetChatID(ctx, obj.ChatID)
}

// Sender is the resolver for the sender field.
func (r *chatDetailsResolver) Sender(ctx context.Context, obj *model.ChatDetails) (*model.User, error) {
	return service.UserGetByID(ctx, obj.SenderID)
}

// CreateNewMessage is the resolver for the CreateNewMessage field.
func (r *mutationResolver) CreateNewMessage(ctx context.Context, message model.NewMessage) (*model.ChatDetails, error) {
	var header *model.ChatHeader
	if r.DB.First(&header).Error != nil {
		//no header found (no chat available)
		header = &model.ChatHeader{ChatID: uuid.NewString()}
		r.DB.Save(&header)
	}

	payloadJwt := middlewares.CtxValue(ctx)

	newMessage := &model.ChatDetails{
		ChatID:    header.ChatID,
		SenderID:  payloadJwt.UserId,
		Message:   message.Message,
		CreatedAt: time.Time{},
	}

	return newMessage, r.DB.Save(&newMessage).Error

}

// GetAllChatData is the resolver for the GetAllChatData field.
func (r *queryResolver) GetAllChatData(ctx context.Context, chatID *string) (*model.ChatDetails, error) {
	panic(fmt.Errorf("not implemented: GetAllChatData - GetAllChatData"))
}

// ChatDetails returns ChatDetailsResolver implementation.
func (r *Resolver) ChatDetails() ChatDetailsResolver { return &chatDetailsResolver{r} }

type chatDetailsResolver struct{ *Resolver }
