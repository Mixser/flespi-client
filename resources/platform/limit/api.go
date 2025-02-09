package flespi_limit

import (
	"fmt"

	"github.com/mixser/flespi-client"
)

func NewLimit(c *flespi.Client, name string, options ...CreateLimitOption) (*Limit, error) {
	limit := Limit{Name: name}

	for _, opt := range options {
		opt(&limit)
	}

	response := limitsResponse{}

	err := c.RequestAPI("POST", "platform/limits", []Limit{limit}, &response)

	if err != nil {
		return nil, err
	}

	return &response.Limits[0], nil
}

func ListLimits(c *flespi.Client) ([]Limit, error) {
	response := limitsResponse{}

	err := c.RequestAPI("GET", "platform/limits/all", nil, &response)

	if err != nil {
		return nil, err
	}

	return response.Limits, nil
}

func GetLimit(c *flespi.Client, limitId int64) (*Limit, error) {
	response := limitsResponse{}

	err := c.RequestAPI("GET", fmt.Sprintf("platform/limits/%d", limitId), nil, &response)

	if err != nil {
		return nil, err
	}

	return &response.Limits[0], nil
}

func UpdateLimit(c *flespi.Client, limit Limit) (*Limit, error) {
	if limit.Id == 0 {
		return nil, fmt.Errorf("ID must be provided")
	}

	limitId := limit.Id
	limit.Id = 0

	response := limitsResponse{}

	err := c.RequestAPI("PUT", fmt.Sprintf("platform/limits/%d", limitId), limit, &response)

	if err != nil {
		return nil, err
	}

	return &response.Limits[0], nil
}


func DeleteLimit(c *flespi.Client, limit Limit) (error) {
	if limit.Id == 0 {
		return fmt.Errorf("ID must be provided")
	}

	return DeleteLimitById(c, limit.Id)
}

func DeleteLimitById(c *flespi.Client, limitId int64) (error) {
	err := c.RequestAPI("DELETE", fmt.Sprintf("platform/limits/%d", limitId), nil, nil)

	if err != nil {
		return err
	}

	return nil
}
