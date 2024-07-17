package controller

import (
	"net/http"
	"strconv"
	"tokyokitten/data/request"
	"tokyokitten/data/response"
	"tokyokitten/helper"
	"tokyokitten/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type KittensController struct {
	kittensService service.KittensService
}

func NewKittensController(service service.KittensService) *KittensController {
	return &KittensController{
		kittensService: service,
	}
}

// Createkittens		godoc
// @Summary			Create kittens
// @Description		Save kittens data in Db.
// @Param			kittens body request.CreatekittensRequest true "Create kittens"
// @Produce			application/json
// @kittens			kittens
// @Success			200 {object} response.Response{}
// @Router			/kittens [post]
func (controller *KittensController) Create(ctx *gin.Context) {
	log.Info().Msg("create kittens")
	createkittensRequest := request.CreateKittensRequest{}
	err := ctx.ShouldBindJSON(&createkittensRequest)
	helper.ErrorPanic(err)

	controller.kittensService.Create(createkittensRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Updatekittens		godoc
// @Summary			Update kittens
// @Description		Update kittens data.
// @Param			tagId path string true "update kittens by id"
// @Param			kittens body request.CreatekittensRequest true  "Update kittens"
// @kittens			kittens
// @Produce			application/json
// @Success			200 {object} response.Response{}
// @Router			/kittens/{tagId} [patch]
func (controller *KittensController) Update(ctx *gin.Context) {
	log.Info().Msg("update kittens")
	updatekittensRequest := request.UpdateKittensRequest{}
	err := ctx.ShouldBindJSON(&updatekittensRequest)
	helper.ErrorPanic(err)

	kittenId := ctx.Param("kittenId")
	id, err := strconv.Atoi(kittenId)
	helper.ErrorPanic(err)
	updatekittensRequest.Id = id

	controller.kittensService.Update(updatekittensRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Deletekittens		godoc
// @Summary			Delete kittens
// @Description		Remove kittens data by id.
// @Produce			application/json
// @kittens			kittens
// @Success			200 {object} response.Response{}
// @Router			/kittens/{tagID} [delete]
func (controller *KittensController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete kittens")
	kittenId := ctx.Param("kittenId")
	id, err := strconv.Atoi(kittenId)
	helper.ErrorPanic(err)
	controller.kittensService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindByIdkittens 		godoc
// @Summary				Get Single kittens by id.
// @Param				tagId path string true "update kittens by id"
// @Description			Return the tahs whoes tagId valu mathes id.
// @Produce				application/json
// @kittens				kittens
// @Success				200 {object} response.Response{}
// @Router				/kittens/{tagId} [get]
func (controller *KittensController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid kittens")
	kittenId := ctx.Param("kittenId")
	id, err := strconv.Atoi(kittenId)
	helper.ErrorPanic(err)

	kittenResponse := controller.kittensService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   kittenResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAllkittens 		godoc
// @Summary			Get All kittens.
// @Description		Return list of kittens.
// @kittens			kittens
// @Success			200 {obejct} response.Response{}
// @Router			/kittens [get]
func (controller *KittensController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll kittens")
	tagResponse := controller.kittensService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
