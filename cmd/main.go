package main

import (
	"flag"
	"image-sync-data/config"
	"image-sync-data/web/common"
	"image-sync-data/web/router"
	"log"
)

func init() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.toml", "Path to the configuration file")
	flag.Parse()
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}
	common.InitializeCommonServices(cfg)

}

func main() {
	r := router.InitRouter()
	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
