package flespi_geofence

import "github.com/mixser/flespi-client/internal/flespiapi"

// GeofenceClient provides receiver-based methods for managing Flespi geofences.
// Access it via Client.Geofences after creating a flespi.Client.
type GeofenceClient struct {
	c flespiapi.APIRequester
}

// NewGeofenceClient creates a GeofenceClient wrapping the given flespiapi.APIRequester.
func NewGeofenceClient(c flespiapi.APIRequester) *GeofenceClient {
	return &GeofenceClient{c: c}
}

func (gc *GeofenceClient) Create(name string, options ...CreateGeofenceOption) (*Geofence, error) {
	return NewGeofence(gc.c, name, options...)
}

func (gc *GeofenceClient) List() ([]Geofence, error) {
	return ListGeofences(gc.c)
}

func (gc *GeofenceClient) GetById(geofenceId int64) (*Geofence, error) {
	return GetGeofence(gc.c, geofenceId)
}

func (gc *GeofenceClient) Update(geofence Geofence) (*Geofence, error) {
	return UpdateGeofence(gc.c, geofence)
}

func (gc *GeofenceClient) Delete(geofence Geofence) error {
	return DeleteGeofence(gc.c, geofence)
}

func (gc *GeofenceClient) DeleteById(geofenceId int64) error {
	return DeleteGeofenceById(gc.c, geofenceId)
}
