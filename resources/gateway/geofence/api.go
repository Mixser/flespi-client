package flespi_geofence

import (
	"fmt"
	"github.com/mixser/flespi-client"
)

func ListGeofences(c *flespi.Client) ([]Geofence, error) {
	response := geofencesResponse{}

	err := c.RequestAPI("GET", "gw/geofences/all?fields=id,geometry", nil, &response)

	if err != nil {
		return nil, err
	}

	return response.Geofences, nil
}

func NewGeofence(c *flespi.Client, name string, options ...CreateGeofenceOption) (*Geofence, error) {
	geofence := Geofence{Name: name}

	for _, opt := range options {
		opt(&geofence)
	}

	response := geofencesResponse{}

	err := c.RequestAPI("POST", "gw/geofences?fields=id,geometry", []Geofence{geofence}, &response)

	if err != nil {
		return nil, err
	}

	return &response.Geofences[0], nil
}

func UpdateGeofence(c *flespi.Client, geofence Geofence) (*Geofence, error) {
	response := geofencesResponse{}

	geofenceId := geofence.Id
	geofence.Id = 0

	err := c.RequestAPI("PUT", fmt.Sprintf("gw/geofences/%d", geofenceId), geofence, &response)

	if err != nil {
		return nil, err
	}

	return &response.Geofences[0], nil
}

func DeleteGeofence(c *flespi.Client, geofence Geofence) error {
	return DeleteGeofenceById(c, geofence.Id)
}

func DeleteGeofenceById(c *flespi.Client, geofenceId int64) error {
	return c.RequestAPI("DELETE", fmt.Sprintf("gw/geofences/%d", geofenceId), nil, nil)
}
