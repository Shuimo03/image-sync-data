package docker

import (
	"context"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"strings"
)

const latestTag = "latest"
const defaultTarget = ""

type DockerClient struct {
	cli *client.Client
}

func (dc *DockerClient) ImageTagList() ([]string, error) {
	ctx := context.Background()
	imagesList, err := dc.cli.ImageList(ctx, image.ListOptions{})
	var name string
	var tag string
	repoTags := make([]string, 0)
	imageInfos := make([]string, 0)
	if err != nil {
		return nil, err
	}
	for _, imageList := range imagesList {
		repoTags = append(repoTags, imageList.RepoTags...)
	}
	for _, repoTag := range repoTags {
		name, tag = parseRepoTag(repoTag)
		imageInfos = append(imageInfos, name+":"+tag)
	}

	return imageInfos, nil
}

// 如果没有指定tag,默认使用latest,默认REPOSITORY为sync-image
func (dc *DockerClient) PushImage(image, repository, tag string) {

}

// TODO 具体业务需要分离
func (dc *DockerClient) selectImages(selectImage string) ([]string, error) {
	imagesList, err := dc.ImageTagList()
	if err != nil {
		return nil, err
	}
	selectedImages := make([]string, 0, len(imagesList))
	for _, imageName := range imagesList {
		if imageName == selectImage {
			selectedImages = append(selectedImages, imageName)
		}
	}
	return selectedImages, nil
}

//func (dc *DockerClient) TagImage(sourceImage, targetImage string) ([]string, error) {
//	if targetImage == "" {
//
//	}
//	images, err := dc.selectImages(sourceImage)
//	tagNames := make([]string, 0, len(images))
//	if err != nil {
//		return nil, err
//	}
//	for _, image := range images {
//
//	}
//}

func parseRepoTag(repoTag string) (string, string) {
	// 找到最后一个冒号的位置
	colonIndex := strings.LastIndex(repoTag, ":")
	if colonIndex == -1 {
		// 如果没有找到冒号，返回整个字符串作为 name，默认 tag 为 "latest"
		return repoTag, latestTag
	}

	// 分割字符串为 name 和 tag
	name := repoTag[:colonIndex]
	tag := repoTag[colonIndex+1:]

	return name, tag
}

func NewDockerClient() (*DockerClient, error) {
	cli, clientError := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if clientError != nil {
		return nil, clientError
	}
	defer cli.Close()
	return &DockerClient{cli: cli}, nil
}
