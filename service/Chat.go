package service

import (
	"context"
	"github.com/liquedgit/tokoMeLia/Database"
	"github.com/liquedgit/tokoMeLia/graph/model"
)

func GetChatID(ctx context.Context, id string) (*model.ChatHeader, error) {
	db := Database.GetInstance()
	var chat *model.ChatHeader
	return chat, db.First(&chat, "chat_id = ?", id).Error
}

func GetAllChatByID(ctx context.Context, chatId string) ([]*model.ChatDetails, error) {
	db := Database.GetInstance()
	var chatData []*model.ChatDetails
	return chatData, db.Find(&chatData, "chat_id = ?", chatId).Error
}
