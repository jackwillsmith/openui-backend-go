package model

import (
	"github.com/openui-backend-go/common/database"
)

var _ ChatModel = (*customChatModel)(nil)

type (
	// ChatModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatModel.
	ChatModel interface {
		chatModel
	}

	customChatModel struct {
		*defaultChatModel
	}
)

// NewChatModel returns a model for the database table.
func NewChatModel(db *database.GormDao, cache *database.DcRedisClient) ChatModel {
	return &customChatModel{
		defaultChatModel: newChatModel(db, cache),
	}
}
