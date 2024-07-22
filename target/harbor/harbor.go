package harbor

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const HarborAPI = "https://192.168.126.3:40443/api/v2.0"

var loginAPI = "https://192.168.126.3:40443/c/login"

type Harbor struct {
	client  *http.Client
	Options Options
}

func CreateRepository(repositoryName string) {

}

func ListRepositories() ([]string, error) {
	panic("implement me")
}

func (h *Harbor) Login() (string, error) {
	requestBody, err := json.Marshal(h.Options)
	if err != nil {
		return "", err
	}
	resp, err := h.client.Post(loginAPI, "application/x-www-form-urlencoded", bytes.NewBuffer(requestBody))
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (h *Harbor) Ping() (string, error) {
	resp, err := h.client.Get(HarborAPI + "/ping")
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

func init() {
}

func NewHarborClient(opts *Options) *Harbor {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //这里作为开发和测试，所以忽略了https证书
		},
	}

	client := &Harbor{
		Options: *opts,
		client:  httpClient,
	}

	return client
}
