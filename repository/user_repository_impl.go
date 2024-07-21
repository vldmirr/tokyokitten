package repository

import (
	"errors"
	"tokyokitten/data/request"
	"tokyokitten/helper"
	"tokyokitten/model"

	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsersRepositoryImpl(Db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{Db: Db}
}

// Delete implements UsersRepository
func (u *UsersRepositoryImpl) Delete(usersId int) {
	var users model.User
	result := u.Db.Where("id = ?", usersId).Delete(&users)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UsersRepository
func (u *UsersRepositoryImpl) FindAll() []model.User {
	var users []model.User
	results := u.Db.Find(&users)
	helper.ErrorPanic(results.Error)
	return users
}

// FindById implements UsersRepository
func (u *UsersRepositoryImpl) FindById(usersId int) (model.User, error) {
	var users model.User
	result := u.Db.Find(&users, usersId)
	if result != nil {
		return users, nil
	} else {
		return users, errors.New("users is not found")
	}
}

// Save implements UsersRepository
func (u *UsersRepositoryImpl) Save(users model.User) {
	result := u.Db.Create(&users)
	helper.ErrorPanic(result.Error)
}

// Update implements UsersRepository
func (u *UsersRepositoryImpl) Update(users model.User) {
	var updateUsers = request.UpdateUsersRequest{
		Id:       users.Id,
		Username: users.Username,
		Email:    users.Email,
		Password: users.Password,
	}
	result := u.Db.Model(&users).Updates(updateUsers)
	helper.ErrorPanic(result.Error)
}

// FindByUsername implements UsersRepository
func (u *UsersRepositoryImpl) FindByUsername(username string) (model.User, error) {
	var users model.User
	result := u.Db.First(&users, "username = ?", username)

	if result.Error != nil {
		return users, errors.New("invalid username or Password")
	}
	return users, nil
}