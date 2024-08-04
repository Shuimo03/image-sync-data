package router

import (
	"github.com/gin-gonic/gin"
	"image-sync-data/web/handlers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.POST("/list/repository", handlers.ListRepositoriesInfo)
	return r
}
