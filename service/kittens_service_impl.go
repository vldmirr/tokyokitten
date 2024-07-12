package service

import (
	
	"log"
	"tokyokitten/data/request"
	"tokyokitten/data/response"
	"tokyokitten/helper"
	"tokyokitten/model"
	"tokyokitten/repository"

	"github.com/go-playground/validator/v10"
)

type KittensServiceImpl struct {
	kittensRepository repository.KittensRepository
	Validate          *validator.Validate
}

func NewKittensServiceImpl(kittenRepository repository.KittensRepository, validate *validator.Validate) KittensService {
	return &KittensServiceImpl{
		kittensRepository: kittenRepository,
		Validate:          validate,
	}
}

// Create implements kittensService
func (t *KittensServiceImpl) Create(kittens request.CreateKittensRequest) {
	err := t.Validate.Struct(kittens)

	

	helper.ErrorPanic(err)
	
	if kittens.Count>0 {
		kittens.InStock=true
	}else{
		kittens.InStock=false
	}

	kittenModel := model.Kitten{
		Name:    kittens.Name,
		Age:     kittens.Age,
		Color:   kittens.Name,
		Breed:   kittens.Breed,
		Price:   kittens.Price,
		Count:   kittens.Count,
		InStock: kittens.InStock,
	}
	log.Printf("ncoming request: %+v\n", kittens)
	t.kittensRepository.Save(kittenModel)
}

// Delete implements kittensService
func (t *KittensServiceImpl) Delete(kittensId int) {
	t.kittensRepository.Delete(kittensId)
}

// FindAll implements kittensService
func (t *KittensServiceImpl) FindAll() []response.Kittens_Response {
	result := t.kittensRepository.FindAll()

	var kittens []response.Kittens_Response
	for _, value := range result {
		kitten := response.Kittens_Response{
			Id:   value.Id,
			Name: value.Name,
		}
		kittens = append(kittens, kitten)
	}

	return kittens
}

// FindById implements kittensService
func (t *KittensServiceImpl) FindById(kittensId int) response.Kittens_Response {
	kittenData, err := t.kittensRepository.FindById(kittensId)
	helper.ErrorPanic(err)

	kittenResponse := response.Kittens_Response{
		Id:   kittenData.Id,
		Name: kittenData.Name,
	}
	return kittenResponse
}

// Update implements kittensService
func (t *KittensServiceImpl) Update(kittens request.UpdateKittensRequest) {
	kittenData, err := t.kittensRepository.FindById(kittens.Id)
	helper.ErrorPanic(err)
	kittenData.Name = kittens.Name
	t.kittensRepository.Update(kittenData)
}
