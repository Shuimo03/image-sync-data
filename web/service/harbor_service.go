package service

import (
	"image-sync-data/target/harbor"
)

type HarborService struct {
	Client *harbor.Harbor
}

func (hs *HarborService) ListRepositoriesInfoService(page, size int) ([]harbor.Repository, error) {
	repositories, err := hs.Client.ListRepositories(page, size)
	if err != nil {
		return nil, err
	}
	return repositories, nil
}

func NewHarborService(opts *harbor.Options) *HarborService {
	client := harbor.NewHarborClient(opts)
	hs := &HarborService{
		Client: client,
	}
	return hs
}
