// Package service contain all business logic
package service

import (
	"github.com/103cuong/grpc_go_kit/configs"
	"github.com/103cuong/grpc_go_kit/pkg/model"
)

// GetCategories fetch all categories
func GetCategories(categories *[]model.Category) (err error) {
	err = configs.DB.Preload("Cats").Find(&categories).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateCategory create new category
func CreateCategory(category *model.Category) (err error) {
	err = configs.DB.Create(&category).Error
	if err != nil {
		return err
	}
	return nil
}
