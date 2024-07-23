package router

import (
	"github.com/gin-gonic/gin"
	"image-sync-data/web/handlers"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Group("/v1")
	r.POST("/list/repository", handlers.ListRepositoriesInfo)
	return r
}
