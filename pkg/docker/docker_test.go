package docker

import (
	"fmt"
	"testing"
)

func TestDockerClient_ImageTagList(t *testing.T) {
	dockerClient, err := NewDockerClient()
	if err != nil {
		t.Fatal(err)
	}
	images, err := dockerClient.ImageTagList()
	if err != nil {
		t.Fatal(err)
	}
	for _, image := range images {
		fmt.Println(image)
	}
}
