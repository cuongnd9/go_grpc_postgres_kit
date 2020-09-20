package service

import (
	"context"
	"github.com/103cuong/go_grpc_postgres_kit/configs"
	"github.com/103cuong/go_grpc_postgres_kit/pkg/model"
)

// GetCats fetch all cats.
func GetCats(ctx context.Context, cats *[]model.Cat) (err error) {
	err = configs.DB.WithContext(ctx).Preload("Category").Find(&cats).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateCat create new cat.
func CreateCat(ctx context.Context, cat *model.Cat) (err error) {
	err = configs.DB.WithContext(ctx).Create(&cat).Error
	if err != nil {
		return err
	}
	return nil
}

// GetCatByID fetch cat by ID.
func GetCatByID(ctx context.Context, cat *model.Cat, id string) (err error) {
	err = configs.DB.WithContext(ctx).Preload("Category").Where("id = ?", id).First(&cat).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateCat update cat.
func UpdateCat(ctx context.Context, cat *model.Cat) (err error) {
	err = configs.DB.WithContext(ctx).Save(&cat).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteCat delete cat.
func DeleteCat(ctx context.Context, cat *model.Cat, id string) (err error) {
	err = configs.DB.WithContext(ctx).Where("id = ?", id).Delete(&cat).Error
	if err != nil {
		return err
	}
	return nil
}
