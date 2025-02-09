package flespi_geofence

import "encoding/json"

func UnmarshalGeometry(rawValue json.RawMessage) (GeofenceGeometry, error) {
	var err error = nil

	var circle Circle

	if err = json.Unmarshal(rawValue, &circle); err == nil {
		return &circle, nil
	}

	var polygon Polygon

	if err = json.Unmarshal(rawValue, &polygon); err == nil {
		return &polygon, nil
	}

	var corridor Corridor

	if err = json.Unmarshal(rawValue, &corridor); err == nil {
		return &corridor, nil
	}

	return nil, err
}
