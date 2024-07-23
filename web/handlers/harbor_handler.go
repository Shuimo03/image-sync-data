package handlers

import (
	"github.com/gin-gonic/gin"
	v1 "image-sync-data/web/api/v1"
	"image-sync-data/web/service"
	"net/http"
)

type HarborHandler struct {
	HS *service.HarborService
}

func (hs *HarborHandler) ListRepositoriesInfo(c *gin.Context) {
	var registry v1.Registry
	if err := c.ShouldBindJSON(&registry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read registry"})
		return
	}
	registries, err := hs.HS.ListRepositoriesInfoService(registry.Page, registry.Size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, registries)
}
