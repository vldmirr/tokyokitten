package router

import (
	"tokyokitten/controller"

	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(kittensController *controller.KittensController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	kittensRouter := baseRouter.Group("/kittens")
	kittensRouter.GET("", kittensController.FindAll)
	kittensRouter.GET("/:kittenId", kittensController.FindById)
	kittensRouter.POST("", kittensController.Create)
	kittensRouter.PATCH("/:kittenId", kittensController.Update)
	kittensRouter.DELETE("/:kittenId", kittensController.Delete)

	return router
}
