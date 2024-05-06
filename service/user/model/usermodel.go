package model

import (
    "openui-backend-go/common/database"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(db *database.GormDao, cache *database.DcRedisClient) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(db, cache),
	}
}
