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
		Username: "admin",
		Password: "Harbor12345",
	}
	harborClient := NewHarborClient(harborOption)
	pong, err := harborClient.Ping()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(pong)
}
