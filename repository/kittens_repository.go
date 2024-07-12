package repository

import "tokyokitten/model"

type KittensRepository interface {
	Save(kitten model.Kitten)
	Update(kittens model.Kitten)
	Delete(kittensId int)
	FindById(kittensId int) (kittens model.Kitten, err error)
	FindAll() []model.Kitten
}
