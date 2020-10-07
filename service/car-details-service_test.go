package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	carDetailsService = NewCarDetailsService()
)

func TestGetDetails(t *testing.T) {

	carDetails := carDetailsService.GetDetails()
	require.NotNil(t, carDetails)
	require.Equal(t, 1, carDetails.ID)
	require.Equal(t, "Montero", carDetails.Model)
	require.Equal(t, 2002, carDetails.Year)
	require.Equal(t, "Alyosha", carDetails.OwnerFirstName)

}
