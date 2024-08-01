package repository

import (
	"errors"
	"tokyokitten/data/request"
	"tokyokitten/helper"
	"tokyokitten/model"
	//"tokyokitten/utils"
	"gorm.io/gorm"
	//"github.com/redis/go-redis/v9"
)

type KittenRepositoryImpl struct {
	Db *gorm.DB
}

func NewKittensRepositoryImpl(Db *gorm.DB) KittensRepository {
	return &KittenRepositoryImpl{Db: Db}
}

// Delete implements kittensRepository
func (t *KittenRepositoryImpl) Delete(kittensId int) {
	var kittens model.Kitten
	result := t.Db.Where("id = ?", kittensId).Delete(&kittens)
	helper.ErrorPanic(result.Error)
}

// FindAll implements kittensRepository
func (t *KittenRepositoryImpl) FindAll() []model.Kitten {
	var kittens []model.Kitten
	result := t.Db.Find(&kittens)
	helper.ErrorPanic(result.Error)
	return kittens
}

// FindById implements kittensRepository
func (t *KittenRepositoryImpl) FindById(kittensId int) (kittens model.Kitten, err error) {
	var kitten model.Kitten
	result := t.Db.Find(&kitten, kittensId)
	if result != nil {
		return kitten, nil
	} else {
		return kitten, errors.New("kitten is not found")
	}
}

// Save implements kittensRepository
func (t *KittenRepositoryImpl) Save(kitten model.Kitten) {
	result := t.Db.Create(&kitten)
	helper.ErrorPanic(result.Error)
}

// Update implements kittensRepository
func (t *KittenRepositoryImpl) Update(kittens model.Kitten) {
	var updateKitten = request.UpdateKittensRequest{
		Id:   kittens.Id,
		Name: kittens.Name,
	}
	result := t.Db.Model(&kittens).Updates(updateKitten)
	helper.ErrorPanic(result.Error)
}
