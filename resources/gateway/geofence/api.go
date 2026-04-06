package flespi_geofence

import (
	"fmt"
	"github.com/mixser/flespi-client/internal/flespiapi"
)

func ListGeofences(c flespiapi.APIRequester) ([]Geofence, error) {
	response := geofencesResponse{}

	err := c.RequestAPI("GET", "gw/geofences/all?fields=id,name,enabled,priority,geometry,cid", nil, &response)

	if err != nil {
		return nil, err
	}

	return response.Geofences, nil
}

func GetGeofence(c flespiapi.APIRequester, geofenceId int64) (*Geofence, error) {
	response := geofencesResponse{}

	err := c.RequestAPI("GET", fmt.Sprintf("gw/geofences/%d?fields=id,name,enabled,priority,geometry,cid", geofenceId), nil, &response)

	if err != nil {
		return nil, err
	}

	return &response.Geofences[0], nil
}

func NewGeofence(c flespiapi.APIRequester, name string, options ...CreateGeofenceOption) (*Geofence, error) {
	geofence := Geofence{Name: name}

	for _, opt := range options {
		opt(&geofence)
	}

	var headers map[string]string
	if geofence.AccountId != 0 {
		headers = map[string]string{
			"x-flespi-cid": fmt.Sprintf("%d", geofence.AccountId),
		}
	}

	accountId := geofence.AccountId
	geofence.AccountId = 0
	defer func() { geofence.AccountId = accountId }()

	response := geofencesResponse{}

	if err := c.RequestAPIWithHeaders("POST", "gw/geofences?fields=id,name,enabled,priority,geometry,cid", headers, []Geofence{geofence}, &response); err != nil {
		return nil, err
	}

	return &response.Geofences[0], nil
}

func UpdateGeofence(c flespiapi.APIRequester, geofence Geofence) (*Geofence, error) {
	response := geofencesResponse{}

	geofenceId := geofence.Id
	accountId := geofence.AccountId

	geofence.Id = 0
	geofence.AccountId = 0

	defer func() {
		geofence.Id = geofenceId
		geofence.AccountId = accountId
	}()

	var headers map[string]string
	if accountId != 0 {
		headers = map[string]string{
			"x-flespi-cid": fmt.Sprintf("%d", accountId),
		}
	}

	if err := c.RequestAPIWithHeaders("PUT", fmt.Sprintf("gw/geofences/%d", geofenceId), headers, geofence, &response); err != nil {
		return nil, err
	}

	return &response.Geofences[0], nil
}

func DeleteGeofence(c flespiapi.APIRequester, geofence Geofence) error {
	return DeleteGeofenceById(c, geofence.Id)
}

func DeleteGeofenceById(c flespiapi.APIRequester, geofenceId int64) error {
	return c.RequestAPI("DELETE", fmt.Sprintf("gw/geofences/%d", geofenceId), nil, nil)
}
