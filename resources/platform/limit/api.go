package flespi_limit

import (
	"fmt"

	"github.com/mixser/flespi-client/internal/flespiapi"
)

func NewLimit(c flespiapi.APIRequester, name string, options ...CreateLimitOption) (*Limit, error) {
	limit := Limit{Name: name}

	for _, opt := range options {
		opt(&limit)
	}

	var headers map[string]string
	if limit.AccountId != 0 {
		headers = map[string]string{
			"x-flespi-cid": fmt.Sprintf("%d", limit.AccountId),
		}
	}

	accountId := limit.AccountId
	limit.AccountId = 0
	defer func() { limit.AccountId = accountId }()

	response := limitsResponse{}

	if err := c.RequestAPIWithHeaders("POST", "platform/limits", headers, []Limit{limit}, &response); err != nil {
		return nil, err
	}

	return &response.Limits[0], nil
}

func ListLimits(c flespiapi.APIRequester) ([]Limit, error) {
	response := limitsResponse{}

	err := c.RequestAPI("GET", "platform/limits/all", nil, &response)

	if err != nil {
		return nil, err
	}

	return response.Limits, nil
}

func GetLimit(c flespiapi.APIRequester, limitId int64) (*Limit, error) {
	response := limitsResponse{}

	err := c.RequestAPI("GET", fmt.Sprintf("platform/limits/%d?fields=id,name,description,blocking_duration,api_calls,api_traffic,channels_count,channel_messages,channel_storage,channel_traffic,channel_connections,containers_count,container_storage,cdns_count,cdn_storage,cdn_traffic,devices_count,device_storage,device_media_traffic,device_media_storage,streams_count,stream_storage,stream_traffic,modems_count,mqtt_sessions,mqtt_messages,mqtt_session_storage,mqtt_retained_storage,mqtt_subscriptions,sms_count,tokens_count,subaccounts_count,limits_count,realms_count,calcs_count,calcs_storage,plugins_count,plugin_traffic,plugin_buffered_messages,groups_count,webhooks_count,webhook_storage,webhook_traffic,grants_count,identity_providers_count,cid", limitId), nil, &response)

	if err != nil {
		return nil, err
	}

	return &response.Limits[0], nil
}

func UpdateLimit(c flespiapi.APIRequester, limit Limit) (*Limit, error) {
	if limit.Id == 0 {
		return nil, fmt.Errorf("ID must be provided")
	}

	limitId := limit.Id
	accountId := limit.AccountId

	limit.Id = 0
	limit.AccountId = 0

	defer func() {
		limit.Id = limitId
		limit.AccountId = accountId
	}()

	var headers map[string]string
	if accountId != 0 {
		headers = map[string]string{
			"x-flespi-cid": fmt.Sprintf("%d", accountId),
		}
	}

	response := limitsResponse{}

	if err := c.RequestAPIWithHeaders("PUT", fmt.Sprintf("platform/limits/%d", limitId), headers, limit, &response); err != nil {
		return nil, err
	}

	return &response.Limits[0], nil
}

func DeleteLimit(c flespiapi.APIRequester, limit Limit) error {
	if limit.Id == 0 {
		return fmt.Errorf("ID must be provided")
	}

	return DeleteLimitById(c, limit.Id)
}

func DeleteLimitById(c flespiapi.APIRequester, limitId int64) error {
	err := c.RequestAPI("DELETE", fmt.Sprintf("platform/limits/%d", limitId), nil, nil)

	if err != nil {
		return err
	}

	return nil
}
