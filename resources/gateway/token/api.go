package flespi_token

import (
	"fmt"
	"github.com/mixser/flespi-client"
)

func NewToken(c *flespi.Client, info string, options ...CreateTokenOption) (*Token, error) {
	token := Token{Info: info}

	for _, opt := range options {
		opt(&token)
	}

	response := tokensResponse{}

	err := c.RequestAPI("POST", "platform/tokens", []Token{token}, &response)

	if err != nil {
		return nil, err
	}


	return &response.Tokens[0], nil
}


func ListTokens(c *flespi.Client) ([]Token, error) {
	response := tokensResponse{}

	err := c.RequestAPI("GET", "platform/tokens/all", nil, &response)

	if err != nil {
		return nil, err
	}

	return response.Tokens, nil
}

func GetToken(c *flespi.Client, tokenId int64) (*Token, error) {
	response := tokensResponse{}

	err := c.RequestAPI("GET", fmt.Sprintf("platform/tokens/%d", tokenId), nil, &response)

	if err != nil {
		return nil, err
	}

	return &response.Tokens[0], nil
}

func UpdateToken(c *flespi.Client, token Token) (*Token, error) {
	response := tokensResponse{}

	tokenId := token.Id
	
	token.Id = 0
	token.Key = ""

	err := c.RequestAPI("PUT", fmt.Sprintf("platform/tokens/%d", tokenId), token, &response)

	if err != nil {
		return nil, err
	}

	return &response.Tokens[0], nil
}

func DeleteToken(c *flespi.Client, token Token) error {
	return DeleteTokenById(c, token.Id)
}

func DeleteTokenById(c *flespi.Client, tokenId int64) error {
	return c.RequestAPI("DELETE", fmt.Sprintf("platform/tokens/%d", tokenId), nil, nil)
}