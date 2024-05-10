package model

import "github.com/openui-backend-go/common/database"

var _ ProductModel = (*customProductModel)(nil)

type (
	// ProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductModel.
	ProductModel interface {
		productModel
	}

	customProductModel struct {
		*defaultProductModel
	}
)

// NewProductModel returns a model for the database table.
func NewProductModel(db *database.GormDao, cache *database.DcRedisClient) ProductModel {
	return &customProductModel{
		defaultProductModel: newProductModel(db, cache),
	}
}
