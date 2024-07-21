package router

import (
	"tokyokitten/controller"
	"tokyokitten/middleware"

	"net/http"
	"tokyokitten/repository"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(userRepository repository.UsersRepository,
	authenticationController *controller.AuthenticationController,
	usersController *controller.UserController,
	kittensController *controller.KittensController) *gin.Engine {
	router := gin.Default()

	//swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hmmmm))))")
	})

	baseRouter := router.Group("/api")

	//auth
	authenticationRouter := baseRouter.Group("/auth")
	authenticationRouter.POST("/register", authenticationController.Register)
	authenticationRouter.POST("/login", authenticationController.Login)

	usersRouter := baseRouter.Group("/users")
	usersRouter.GET("", middleware.DeserializeUser(userRepository), usersController.GetUsers)
	//kitten
	kittensRouter := baseRouter.Group("/kittens")
	kittensRouter.GET("", kittensController.FindAll)
	kittensRouter.GET("/:kittenId", kittensController.FindById)
	kittensRouter.POST("", kittensController.Create)
	kittensRouter.PATCH("/:kittenId", kittensController.Update)
	kittensRouter.DELETE("/:kittenId", kittensController.Delete)

	return router
}
