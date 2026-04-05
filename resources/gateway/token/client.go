package flespi_token

import "github.com/mixser/flespi-client/internal/flespiapi"

// TokenClient provides receiver-based methods for managing Flespi tokens.
// Access it via Client.Tokens after creating a flespi.Client.
type TokenClient struct {
	c flespiapi.APIRequester
}

// NewTokenClient creates a TokenClient wrapping the given flespiapi.APIRequester.
func NewTokenClient(c flespiapi.APIRequester) *TokenClient {
	return &TokenClient{c: c}
}

func (tc *TokenClient) Create(info string, options ...CreateTokenOption) (*Token, error) {
	return NewToken(tc.c, info, options...)
}

func (tc *TokenClient) List() ([]Token, error) {
	return ListTokens(tc.c)
}

func (tc *TokenClient) Get(tokenId int64) (*Token, error) {
	return GetToken(tc.c, tokenId)
}

func (tc *TokenClient) Update(token Token) (*Token, error) {
	return UpdateToken(tc.c, token)
}

func (tc *TokenClient) Delete(token Token) error {
	return DeleteToken(tc.c, token)
}

func (tc *TokenClient) DeleteById(tokenId int64) error {
	return DeleteTokenById(tc.c, tokenId)
}
