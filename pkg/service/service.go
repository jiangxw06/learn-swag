package service

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/files" // swagger embed files
)

func InitRouters(group *gin.RouterGroup) {
	group.Handle("GET", "/users/:user", UserService.GetUser)
}

// @title                      Learn Swag
// @version                    1.0
// @securityDefinitions.basic  BasicAuth
func RunSvc() {
	engine := gin.New()
	InitRouters(engine.Group("/api/v1"))
	//engine.GET("/api/v1/users/:user", service.UserService.GetUser)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.Run(":9000")
}