package handlers

import (
	"github.com/gin-gonic/gin"
	v1 "image-sync-data/web/request/v1"
	"image-sync-data/web/service"
	"net/http"
)

func ListRepositoriesInfo(c *gin.Context) {
	var registry v1.Registry
	hs := service.HarborService{}
	if err := c.ShouldBindJSON(&registry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read registry"})
		return
	}
	registries, err := hs.ListRepositoriesInfoService(registry.Page, registry.Size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, registries)
}
