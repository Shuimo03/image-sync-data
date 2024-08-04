package service

import (
	"image-sync-data/pkg/harbor"
	"image-sync-data/web/common"
)

type HarborService struct {
}

func (hs *HarborService) ListRepositoriesInfoService(page int, size int) ([]harbor.Repository, error) {
	repositories, err := common.HarborClient.ListRepositories(page, size)
	if err != nil {
		return nil, err
	}
	return repositories, nil
}
