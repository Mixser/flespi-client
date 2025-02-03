package flespi_cdn

import (
	"fmt"
	"github.com/mixser/flespi-client"
)

func NewCDN(client *flespi.Client, name string, opts ...CreateCDNOption) (*CDN, error) {
	cdn := CDN{Name: name}

	for _, opt := range opts {
		opt(&cdn)
	}

	response := cdnsResponse{}

	err := client.RequestAPI("POST", "storage/cdns", []CDN{cdn}, &response)

	if err != nil {
		return nil, err
	}

	return &response.CDNS[0], nil
}

func ListCDNs(client *flespi.Client) ([]CDN, error) {
	response := cdnsResponse{}

	err := client.RequestAPI("GET", "storage/cdns/all", nil, &response)

	if err != nil {
		return nil, err
	}

	return response.CDNS, nil
}

func GetCDN(client *flespi.Client, cdnId int64) (*CDN, error) {
	response := cdnsResponse{}

	err := client.RequestAPI("GET", fmt.Sprintf("storage/cdns/%d", cdnId), nil, &response)

	if err != nil {
		return nil, err
	}

	return &response.CDNS[0], nil
}

func UpdateCDN(client *flespi.Client, cdn CDN) (*CDN, error) {
	if cdn.Id == 0 {
		return nil, fmt.Errorf("ID must be provided")
	}

	response := cdnsResponse{}
	err := client.RequestAPI("PUT", fmt.Sprintf("storage/cdns/%d", cdn.Id), cdn, &response)

	if err != nil {
		return nil, err
	}

	return &response.CDNS[0], nil
}

func DeleteCDN(client *flespi.Client, cdn CDN) (error) {
	if cdn.Id == 0 {
		return fmt.Errorf("ID must be provided")
	}

	err := DeleteCDNById(client, cdn.Id)

	if err != nil {
		return err
	}

	return nil
}

func DeleteCDNById(client *flespi.Client, cdnId int64) error {
	err := client.RequestAPI("DELETE", fmt.Sprintf("storage/cdns/%d", cdnId), nil, nil)

	if err != nil {
		return err
	}

	return nil
}