package common

import (
	"image-sync-data/config"
	"image-sync-data/pkg/harbor"
)

var HarborClient *harbor.Harbor

func InitializeCommonServices(cfg *config.Config) {
	opts := &harbor.Options{
		Username:        cfg.Target.User,
		Password:        cfg.Target.Password,
		HarborAPIServer: cfg.Target.Server,
	}
	hc := harbor.NewHarborClient(opts)
	HarborClient = hc
}
