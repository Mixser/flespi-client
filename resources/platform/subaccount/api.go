package flespi_subaccount

import (
	"fmt"

	"github.com/mixser/flespi-client/internal/flespiapi"
)

func NewSubaccount(client flespiapi.APIRequester, name string, options ...CreateSubaccountOption) (*Subaccount, error) {
	subaccount := Subaccount{Name: name}

	for _, opt := range options {
		opt(&subaccount)
	}

	var headers map[string]string
	if subaccount.AccountId != 0 {
		headers = map[string]string{
			"x-flespi-cid": fmt.Sprintf("%d", subaccount.AccountId),
		}
	}

	accountId := subaccount.AccountId
	subaccount.AccountId = 0
	defer func() { subaccount.AccountId = accountId }()

	response := subaccountsResponse{}

	if err := client.RequestAPIWithHeaders("POST", "platform/subaccounts", headers, []Subaccount{subaccount}, &response); err != nil {
		return nil, err
	}

	return &response.Subaccounts[0], nil

}

func ListSubaccounts(client flespiapi.APIRequester) ([]Subaccount, error) {
	response := subaccountsResponse{}

	err := client.RequestAPI("GET", "platform/subaccounts/all", nil, &response)

	if err != nil {
		return nil, err
	}

	return response.Subaccounts, nil
}

func GetSubaccount(client flespiapi.APIRequester, subaccountId int64) (*Subaccount, error) {
	response := subaccountsResponse{}

	err := client.RequestAPI("GET", fmt.Sprintf("platform/subaccounts/%d?fields=id,name,limit_id,metadata,cid", subaccountId), nil, &response)

	if err != nil {
		return nil, err
	}

	return &response.Subaccounts[0], nil
}

func UpdateSubaccount(client flespiapi.APIRequester, subaccount Subaccount) (*Subaccount, error) {
	if subaccount.Id == 0 {
		return nil, fmt.Errorf("id should be defined before update")
	}

	subaccountId := subaccount.Id
	accountId := subaccount.AccountId

	subaccount.Id = 0
	subaccount.AccountId = 0

	defer func() {
		subaccount.Id = subaccountId
		subaccount.AccountId = accountId
	}()

	var headers map[string]string
	if accountId != 0 {
		headers = map[string]string{
			"x-flespi-cid": fmt.Sprintf("%d", accountId),
		}
	}

	response := subaccountsResponse{}

	if err := client.RequestAPIWithHeaders("PUT", fmt.Sprintf("platform/subaccounts/%d", subaccountId), headers, subaccount, &response); err != nil {
		return nil, err
	}

	return &response.Subaccounts[0], nil
}

func DeleteSubaccount(client flespiapi.APIRequester, subaccount Subaccount) error {
	if subaccount.Id == 0 {
		return fmt.Errorf("id should be defined before delete")
	}

	return DeleteSubaccountById(client, subaccount.Id)
}

func DeleteSubaccountById(client flespiapi.APIRequester, subaccountId int64) error {
	err := client.RequestAPI("DELETE", fmt.Sprintf("platform/subaccounts/%d", subaccountId), nil, nil)

	if err != nil {
		return err
	}

	return nil
}
