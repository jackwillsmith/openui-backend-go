package database

import (
	"context"
	"fmt"
)

// Create insert data into table
func (d *GormDao) Create(ctx context.Context, tableName string, in interface{}) error {
	// todo: create one record
	if err := d.GetDB().WithContext(ctx).Table(tableName).Create(in).Error; err != nil {
		return err
	}
	return nil
}

// Delete delete table data
func (d *GormDao) Delete(ctx context.Context, tableName string, in interface{}, unscoped bool) error {
	// todo: delete one record
	if unscoped {
		if err := d.GetDB().WithContext(ctx).Table(tableName).Unscoped().Delete(in).Error; err != nil {
			return err
		}
	} else {
		if err := d.GetDB().WithContext(ctx).Table(tableName).Delete(in).Error; err != nil {
			return err
		}
	}
	return nil
}

// DeleteIds delete table multiple data
func (d *GormDao) DeleteIds(ctx context.Context, tableName, column string, ids []int64, in interface{}, unscoped bool) error {
	colName := fmt.Sprintf("`%s` in ?", column)
	if unscoped {
		if err := d.GetDB().WithContext(ctx).Table(tableName).Unscoped().Where(colName, ids).Delete(in).Error; err != nil {
			return err
		}
	} else {
		if err := d.GetDB().WithContext(ctx).Table(tableName).Where(colName, ids).Delete(in).Error; err != nil {
			return err
		}
	}
	return nil
}

// First query one piece of data from table
func (d *GormDao) First(ctx context.Context, tableName string, in interface{}, query interface{}, args ...interface{}) error {
	// todo: query one record
	if err := d.GetDB().WithContext(ctx).Table(tableName).Where(query, args...).First(in).Error; err != nil {
		return err
	}
	return nil
}

// Find query multiple data from table
func (d *GormDao) Find(ctx context.Context, tableName string, in interface{}, query interface{}, args ...interface{}) error {
	// todo: query many records
	if err := d.GetDB().Table(tableName).WithContext(ctx).Where(query, args...).Find(in).Error; err != nil {
		return err
	}
	return nil
}

// Update update one table data
func (d *GormDao) Update(ctx context.Context, tableName string, id uint, updates interface{}) error {
	// todo: update one record
	if err := d.GetDB().Table(tableName).WithContext(ctx).Where("id = ?", id).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

// UpdateOneByParamModel 更新单条数据，根据参数model
func (d *GormDao) UpdateOneByParamModel(ctx context.Context, tableName string, paramModel, updatesModel interface{}) error {
	if err := d.GetDB().Table(tableName).WithContext(ctx).Where(paramModel).Updates(updatesModel).Error; err != nil {
		return err
	}
	return nil
}

func (d *GormDao) Count(ctx context.Context, tableName string, query interface{}, args ...interface{}) (int64, error) {
	var total int64
	if err := d.GetDB().Table(tableName).WithContext(ctx).Where(query, args...).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (d *GormDao) FindByPage(ctx context.Context, tableName string, limit, offset int, in, query interface{}, args ...interface{}) error {
	if err := d.GetDB().Table(tableName).WithContext(ctx).Where(query, args...).Limit(limit).Offset(offset).Find(in).Error; err != nil {
		return err
	}
	return nil
}
