package service

import (
	"tokyokitten/data/request"
	"tokyokitten/data/response"
)

type KittensService interface {
	Create(kittens request.CreateKittensRequest)
	Update(kittens request.UpdateKittensRequest)
	Delete(kittensId int)
	FindById(kittensId int) response.Kittens_Response
	FindAll() []response.Kittens_Response
}
