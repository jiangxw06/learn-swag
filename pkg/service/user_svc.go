package service

import (
	"github.com/gin-gonic/gin"
	"github.com/test/learn-swag/pkg/api"
	"net/http"
)

type (
	userService struct {}
)

// @Summary  Get User
// @description this is the description about getUser api
// @description this is the description2 about getUser api
// @Produce  json
// @Param    user  path      string  true  "User Name"
// @Param    isTest query    boolean true  "Is Test" default(false)
// @Success  200   {object}  api.GetUserResponse
// @Header   all   {string}  Token "token1"
// @Router   /api/v1/users/{user} [get]
func (us userService) GetUser(ctx *gin.Context) {
	name := ctx.Param("user")
	isTest := ctx.Query("isTest")
	msg := "ok"
	if isTest == "true" {
		msg = "test"
	}
	ctx.JSON(http.StatusOK, api.GetUserResponse{
		Code: http.StatusOK,
		Msg:  msg,
		User: us.getUser(name),
	})
}

var (
	UserService = new(userService)
)

func (us userService) getUser(name string) api.User {
	return api.User{
		ID: 1,
		Name: name,
		Gender: api.Male,
		Age: 12,
	}
}