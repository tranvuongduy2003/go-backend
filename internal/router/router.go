package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tranvuongduy2003/go-backend/internal/controller"
	"github.com/tranvuongduy2003/go-backend/internal/middlewares"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.AuthMiddleware())

	v1 := r.Group("/v1/2025")
	{
		v1.GET("/ping/:name", controller.NewPongController().Pong)
		v1.GET("/user/1", controller.NewUserController().GetUserById)
	}

	return r
}