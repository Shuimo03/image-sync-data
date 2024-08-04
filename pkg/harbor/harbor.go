package harbor

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"image-sync-data/pkg/harbor/response"
	"io"
	"net/http"
)

type Harbor struct {
	Client  *http.Client
	Options Options
}

func (h *Harbor) CreateRepository(repositoryName string) {

}

func (h *Harbor) ListRepositories(page, size int) ([]response.Repository, error) {

	url := fmt.Sprintf("%s/repositories?page=%d&page_size=%d", h.Options.HarborAPIServer, page, size)
	resp, err := h.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("failed to fetch repositories: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))
	var repositories []response.Repository
	if jsonError := json.Unmarshal(body, &repositories); jsonError != nil {
		return nil, jsonError
	}

	return repositories, nil
}

func (h *Harbor) ListProjects(page, size int, with_detail bool) ([]response.Project, error) {
	panic("")
}

func (h *Harbor) login() (string, error) {
	return "", nil
}

func (h *Harbor) Ping() (string, error) {
	resp, err := h.Client.Get(h.Options.HarborAPIServer + "/ping")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func NewHarborClient(opts *Options) *Harbor {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //这里作为开发和测试，所以忽略了https证书
		},
	}

	client := &Harbor{
		Options: *opts,
		Client:  httpClient,
	}

	return client
}
