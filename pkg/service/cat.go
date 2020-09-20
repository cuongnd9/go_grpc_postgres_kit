package service

import (
	"github.com/103cuong/grpc_go_kit/configs"
	"github.com/103cuong/grpc_go_kit/pkg/model"
)

// GetCats fetch all cats.
func GetCats(cats *[]model.Cat) (err error) {
	err = configs.DB.Preload("Category").Find(&cats).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateCat create new cat.
func CreateCat(cat *model.Cat) (err error) {
	err = configs.DB.Create(&cat).Error
	if err != nil {
		return err
	}
	return nil
}

// GetCatByID fetch cat by ID.
func GetCatByID(cat *model.Cat, id string) (err error) {
	err = configs.DB.Preload("Category").Where("id = ?", id).First(&cat).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateCat update cat.
func UpdateCat(cat *model.Cat) (err error) {
	err = configs.DB.Save(&cat).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteCat delete cat.
func DeleteCat(cat *model.Cat, id string) (err error) {
	err = configs.DB.Where("id = ?", id).Delete(&cat).Error
	if err != nil {
		return err
	}
	return nil
}
