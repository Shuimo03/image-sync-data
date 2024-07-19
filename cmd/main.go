package main

import "sync-image-data/source/docker"

func main() {
	dc, err := docker.NewDockerClient()
	if err != nil {
		panic(err)
	}
	dc.TagImage("", "")
}
