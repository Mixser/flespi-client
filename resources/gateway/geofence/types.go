package flespi_geofence

import "encoding/json"

type GeofenceGeometry interface {
	GetType() string
}


type Geofence struct {
	Id int64 `json:"id,omitempty"`
	Name string `json:"name"`

	Enabled bool `json:"enabled"`
	Priority int64 `json:"priority"`

	Geometry GeofenceGeometry `json:"geometry"`
}

func (g *Geofence) UnmarshalJSON(data []byte) error {
	var raw struct {
		Id int64 `json:"id"`
		Name string `json:"name"`

		Enabled bool `json:"enabled"`
		Priority int64 `json:"priority"`

		Geometry json.RawMessage `json:"geometry"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	g.Id = raw.Id
	g.Name = raw.Name
	g.Enabled = raw.Enabled
	g.Priority = raw.Priority

	geometry, err := unmarshalGeometry(raw.Geometry)

	if err != nil {
		return err
	}

	g.Geometry = geometry

	return nil
}


type Point struct {
	Latitude float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}


type Circle struct {
	Type   string `json:"type"`
	Center Point  `json:"center"`
	Radius float64 `json:"radius"`
}

func NewCircle(center Point, radius float64) *Circle {
	return &Circle{Type: "circle", Center: center, Radius: radius}
}

func (c *Circle) GetType() string {
	return "circle"
}

type Polygon struct {
	Type string  `json:"type"`
	Path []Point `json:"path"`
}

func NewPolygon(path []Point) *Polygon {
	return &Polygon{Type: "polygon", Path: path}
}

func (p *Polygon) GetType() string {
	return "polygon"
}

type Corridor struct {
	Type string `json:"type"`
	Path []Point `json:"path"`
	Width float64 `json:"width"`
}

func NewCorridor(path []Point, width float64) *Corridor {
	return &Corridor{Type: "corridor", Path: path, Width: width}
}

func (p *Corridor) GetType() string {
	return "corridor"
}


type geofencesResponse struct {
	Geofences []Geofence `json:"result"`
}

type CreateGeofenceOption func(*Geofence)

func WithStatus(enabled bool) CreateGeofenceOption {
	return func(g *Geofence) {
		g.Enabled = enabled
	}
}

func WithPriority(priority int64) CreateGeofenceOption {
	return func(g *Geofence) {
		g.Priority = priority
	}
}

func WithGeometry(geometry GeofenceGeometry) CreateGeofenceOption {
	return func (g *Geofence) {
		g.Geometry = geometry
	}
}