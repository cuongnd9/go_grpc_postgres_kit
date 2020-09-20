// Package service contain all business logic
package service

import (
	"context"
	"github.com/103cuong/go_grpc_postgres_kit/configs"
	"github.com/103cuong/go_grpc_postgres_kit/pkg/model"
)

// GetCategories fetch all categories
func GetCategories(ctx context.Context, categories *[]model.Category) (err error) {
	err = configs.DB.WithContext(ctx).Preload("Cats").Find(&categories).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateCategory create new category
func CreateCategory(ctx context.Context, category *model.Category) (err error) {
	err = configs.DB.WithContext(ctx).Create(&category).Error
	if err != nil {
		return err
	}
	return nil
}
