package flespi_geofence

import (
	"fmt"
	"github.com/mixser/flespi-client/internal/flespiapi"
)

func ListGeofences(c flespiapi.APIRequester) ([]Geofence, error) {
	response := geofencesResponse{}

	err := c.RequestAPI("GET", "gw/geofences/all?fields=id,name,enabled,priority,geometry", nil, &response)

	if err != nil {
		return nil, err
	}

	return response.Geofences, nil
}

func NewGeofence(c flespiapi.APIRequester, name string, options ...CreateGeofenceOption) (*Geofence, error) {
	geofence := Geofence{Name: name}

	for _, opt := range options {
		opt(&geofence)
	}

	response := geofencesResponse{}

	err := c.RequestAPI("POST", "gw/geofences?fields=id,name,enabled,priority,geometry", []Geofence{geofence}, &response)

	if err != nil {
		return nil, err
	}

	return &response.Geofences[0], nil
}

func UpdateGeofence(c flespiapi.APIRequester, geofence Geofence) (*Geofence, error) {
	response := geofencesResponse{}

	geofenceId := geofence.Id
	geofence.Id = 0

	err := c.RequestAPI("PUT", fmt.Sprintf("gw/geofences/%d", geofenceId), geofence, &response)

	if err != nil {
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
