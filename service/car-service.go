package service

import (
	"fmt"
	"net/http"
)

type CarService interface {
	FetchData()
}

const (
	carServiceUrl = "https://myfakeapi.com/api/cars/1"
)

type fetchCarDataService struct{}

func NewCarService() CarService {
	return &fetchCarDataService{}
}

func (*fetchCarDataService) FetchData() {
	client := http.Client{}
	fmt.Printf("Fetching the url %s", carServiceUrl)

	// CHIAMATA API ESTERNA
	resp, err := client.Get(carServiceUrl)
	if err != nil {
		fmt.Errorf("Errore ingaggiando api %v : %v", carServiceUrl, err)
	}

	// RESTITUISCO I DATI AL CANALE
	carDataChannel <- resp
}
