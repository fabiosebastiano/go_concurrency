package service

import (
	"fmt"
	"net/http"
)

type OwnerService interface {
	FetchData()
}

const (
	ownerServiceUrl = "https://myfakeapi.com/api/users/1"
)

type fetchOwnerDataService struct{}

func NewOwnerService() OwnerService {
	return &fetchOwnerDataService{}
}

func (*fetchOwnerDataService) FetchData() {
	client := http.Client{}
	fmt.Printf("Fetching the url %s", ownerServiceUrl)

	// CHIAMATA API ESTERNA
	resp, err := client.Get(ownerServiceUrl)
	if err != nil {
		fmt.Errorf("Errore ingaggiando api %v : %v", ownerServiceUrl, err)
	}

	// RESTITUISCO I DATI AL CANALE
	ownerDataChannel <- resp
}
