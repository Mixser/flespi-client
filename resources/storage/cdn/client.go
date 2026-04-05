package flespi_cdn

import "github.com/mixser/flespi-client/internal/flespiapi"

// CDNClient provides receiver-based methods for managing Flespi CDNs.
// Access it via Client.CDNs after creating a flespi.Client.
type CDNClient struct {
	c flespiapi.APIRequester
}

// NewCDNClient creates a CDNClient wrapping the given flespiapi.APIRequester.
func NewCDNClient(c flespiapi.APIRequester) *CDNClient {
	return &CDNClient{c: c}
}

func (cc *CDNClient) Create(name string, options ...CreateCDNOption) (*CDN, error) {
	return NewCDN(cc.c, name, options...)
}

func (cc *CDNClient) List() ([]CDN, error) {
	return ListCDNs(cc.c)
}

func (cc *CDNClient) Get(cdnId int64) (*CDN, error) {
	return GetCDN(cc.c, cdnId)
}

func (cc *CDNClient) Update(cdn CDN) (*CDN, error) {
	return UpdateCDN(cc.c, cdn)
}

func (cc *CDNClient) Delete(cdn CDN) error {
	return DeleteCDN(cc.c, cdn)
}

func (cc *CDNClient) DeleteById(cdnId int64) error {
	return DeleteCDNById(cc.c, cdnId)
}
