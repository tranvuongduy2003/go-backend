package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tranvuongduy2003/go-backend/internal/service"
	"github.com/tranvuongduy2003/go-backend/response"
)

type UserController struct{
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	response.SuccessResponse(c, 20001, []string{"user1", "user2", "user3"})
	// response.ErrorResponse(c, 20001, "This is a test error message")
}