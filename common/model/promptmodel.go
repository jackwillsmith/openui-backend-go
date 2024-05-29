package model

import (
    "github.com/openui-backend-go/common/database"
)

var _ PromptModel = (*customPromptModel)(nil)

type (
	// PromptModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPromptModel.
	PromptModel interface {
		promptModel
	}

	customPromptModel struct {
		*defaultPromptModel
	}
)

// NewPromptModel returns a model for the database table.
func NewPromptModel(db *database.GormDao, cache *database.DcRedisClient) PromptModel {
	return &customPromptModel{
		defaultPromptModel: newPromptModel(db, cache),
	}
}
