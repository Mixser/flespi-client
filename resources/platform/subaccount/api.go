package flespi_subaccount

import (
	"fmt"

	"github.com/mixser/flespi-client"
)

func NewSubaccount(client *flespi.Client, name string, options ...CreateSubaccountOption) (*Subaccount, error) {
	subaccount := Subaccount{Name: name}

	for _, opt := range options {
		opt(&subaccount)
	}

	response := subaccountsResponse{}

	err := client.RequestAPI("POST", "platform/subaccounts", []Subaccount{subaccount}, &response)

	if err != nil {
		return nil, err
	}

	return &response.Subaccounts[0], nil

}

func ListSubaccounts(client *flespi.Client) ([]Subaccount, error) {
	response := subaccountsResponse{}

	err := client.RequestAPI("GET", "platform/subaccounts/all", nil, &response)

	if err != nil {
		return nil, err
	}

	return response.Subaccounts, nil
}

func GetSubaccount(client *flespi.Client, subaccountId int64) (*Subaccount, error) {
	response := subaccountsResponse{}

	err := client.RequestAPI("GET", fmt.Sprintf("platform/subaccounts/%d", subaccountId), nil, &response)

	if err != nil {
		return nil, err
	}

	return &response.Subaccounts[0], nil
}

func UpdateSubaccount(client *flespi.Client, subaccount Subaccount) (*Subaccount, error) {
	if subaccount.Id == 0 {
		return nil, fmt.Errorf("Id should be defined before update!")
	}

	response := subaccountsResponse{}

	err := client.RequestAPI("PUT", fmt.Sprintf("platform/subaccounts/%d", subaccount.Id), subaccount, &response)

	if err != nil {
		return nil, err
	}

	return &response.Subaccounts[0], nil
}

func DeleteSubaccount(client *flespi.Client, subaccount Subaccount) error {
	if subaccount.Id == 0 {
		return fmt.Errorf("Id should be defined before delete!")
	}

	return DeleteSubaccountById(client, subaccount.Id)
}

func DeleteSubaccountById(client *flespi.Client, subaccountId int64) error {
	err := client.RequestAPI("DELETE", fmt.Sprintf("platform/subaccounts/%d", subaccountId), nil, nil)

	if err != nil {
		return err
	}

	return nil
}
