package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fabiosebastiano/go_concurrency/entity"
)

// CarDetailsService interface
type CarDetailsService interface {
	GetDetails() entity.CarDetails
}

type service struct{}

var (
	carService       CarService   = NewCarService()
	ownerService     OwnerService = NewOwnerService()
	carDataChannel                = make(chan *http.Response)
	ownerDataChannel              = make(chan *http.Response)
)

// NewCarDetailsService  inizializza l'interfaccia e restituisce il servizio
func NewCarDetailsService() CarDetailsService {
	return &service{}
}

func (*service) GetDetails() entity.CarDetails {

	// goroutine che chiama endpoint 1 - ASYNC
	go carService.FetchData()

	// goroutine che chiama endpoint 2 - ASYNC
	go ownerService.FetchData()

	// crea carChannel per i dati
	car, _ := getCarData()

	// crea carOwnerChannel per i dati
	owner, _ := getOwnerData()
	fmt.Println(owner.FirstName)
	fmt.Println(owner.LastName)
	fmt.Println(owner.Email)

	return entity.CarDetails{
		ID:             car.ID,
		Brand:          car.Brand,
		Model:          car.Model,
		Year:           car.Year,
		OwnerFirstName: owner.FirstName,
		OwnerLastName:  owner.LastName,
		OwnerEmail:     owner.Email,
		OwnerJobTitle:  owner.JobTitle,
	}
}

func getCarData() (entity.Car, error) {
	r1 := <-carDataChannel
	var car entity.Car
	err := json.NewDecoder(r1.Body).Decode(&car)
	if err != nil {
		fmt.Println(err.Error())
		return car, err
	}
	return car, nil
}

func getOwnerData() (entity.Owner, error) {
	r1 := <-ownerDataChannel
	var owner entity.Owner
	err := json.NewDecoder(r1.Body).Decode(&owner)
	if err != nil {
		fmt.Println(err.Error())
		return owner, err
	}
	return owner, nil
}
