package repository

import "tokyokitten/model"

type KittensRepository interface {
	Save(kitten model.Kitten)
	Update(kittens model.Kitten)
	Delete(kittensId int)
	FindById(kittensId int) (kittens model.Kitten, err error)
	FindAll() []model.Kitten
}

type UsersRepository interface {
	Save(users model.User)
	Update(users model.User)
	Delete(usersId int)
	FindById(usersId int) (model.User, error)
	FindAll() []model.User
	FindByUsername(username string) (model.User, error)
}
