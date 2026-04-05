package flespi_token

import (
	"fmt"

	"github.com/mixser/flespi-client/internal/flespiapi"
)

func NewToken(c flespiapi.APIRequester, info string, options ...CreateTokenOption) (*Token, error) {
	token := Token{Info: info}

	for _, opt := range options {
		opt(&token)
	}

	accountId := token.AccountId
	token.AccountId = 0

	// Restore AccountId after request
	defer func() { token.AccountId = accountId }()

	var headers map[string]string
	if accountId != 0 {
		headers = map[string]string{
			"x-flespi-cid": fmt.Sprintf("%d", accountId),
		}
	}

	response := tokensResponse{}

	if err := c.RequestAPIWithHeaders("POST", "platform/tokens", headers, []Token{token}, &response); err != nil {
		return nil, err
	}

	return &response.Tokens[0], nil
}

func ListTokens(c flespiapi.APIRequester) ([]Token, error) {
	response := tokensResponse{}

	err := c.RequestAPI("GET", "platform/tokens/all", nil, &response)

	if err != nil {
		return nil, err
	}

	return response.Tokens, nil
}

func GetToken(c flespiapi.APIRequester, tokenId int64) (*Token, error) {
	response := tokensResponse{}

	err := c.RequestAPI("GET", fmt.Sprintf("platform/tokens/%d", tokenId), nil, &response)

	if err != nil {
		return nil, err
	}

	return &response.Tokens[0], nil
}

func UpdateToken(c flespiapi.APIRequester, token Token) (*Token, error) {
	response := tokensResponse{}

	tokenId := token.Id
	tokenKey := token.Key
	tokenAccountId := token.AccountId

	token.Id = 0
	token.Key = ""
	token.AccountId = 0

	defer func() {
		token.Id = tokenId
		token.Key = tokenKey
		token.AccountId = tokenAccountId
	}()

	err := c.RequestAPI("PUT", fmt.Sprintf("platform/tokens/%d", tokenId), token, &response)

	if err != nil {
		return nil, err
	}

	return &response.Tokens[0], nil
}

func DeleteToken(c flespiapi.APIRequester, token Token) error {
	return DeleteTokenById(c, token.Id)
}

func DeleteTokenById(c flespiapi.APIRequester, tokenId int64) error {
	return c.RequestAPI("DELETE", fmt.Sprintf("platform/tokens/%d", tokenId), nil, nil)
}
