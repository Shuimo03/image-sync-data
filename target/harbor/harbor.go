package harbor

import "net/http"

type Harbor struct {
	client http.Client
}

func CreateRepository(repositoryName string) {

}

func ListRepositories() ([]string, error) {
	panic("implement me")
}

func Login(username string, password string) error {
	panic("implement me")
}

func NewHarborClient() (Harbor, error) {
	client := http.Client{}
	return Harbor{client: client}, nil
}
