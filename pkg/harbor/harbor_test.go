package harbor

import (
	"fmt"
	"testing"
)

//func TestHarbor_BasicAuth(t *testing.T) {
//	harborOption := &Options{
//		"admin",
//		"Harbor12345",
//	}
//	harborClient := NewHarborClient(harborOption)
//	harborClient.Ping()
//}

func TestHarbor_Ping(t *testing.T) {
	harborOption := &Options{
		Username:        "admin",
		Password:        "Harbor12345",
		HarborAPIServer: "https://192.168.0.103:40443/api/v2.0",
	}
	harborClient := NewHarborClient(harborOption)
	pong, err := harborClient.Ping()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(pong)
}

func TestHarbor_ListRepositories(t *testing.T) {
	harborOption := &Options{
		Username: "admin",
		Password: "Harbor12345",
	}
	harborClient := NewHarborClient(harborOption)
	repositorys, err := harborClient.ListRepositories(1, 10)
	if err != nil {
		t.Fatal(err)
	}
	for _, repository := range repositorys {
		fmt.Println(repository.Name)
	}
}
