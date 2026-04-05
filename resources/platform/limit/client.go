package flespi_limit

import "github.com/mixser/flespi-client/internal/flespiapi"

// LimitClient provides receiver-based methods for managing Flespi limits.
// Access it via Client.Limits after creating a flespi.Client.
type LimitClient struct {
	c flespiapi.Doer
}

// NewLimitClient creates a LimitClient wrapping the given flespiapi.Doer.
func NewLimitClient(c flespiapi.Doer) *LimitClient {
	return &LimitClient{c: c}
}

func (lc *LimitClient) Create(name string, options ...CreateLimitOption) (*Limit, error) {
	return NewLimit(lc.c, name, options...)
}

func (lc *LimitClient) List() ([]Limit, error) {
	return ListLimits(lc.c)
}

func (lc *LimitClient) Get(limitId int64) (*Limit, error) {
	return GetLimit(lc.c, limitId)
}

func (lc *LimitClient) Update(limit Limit) (*Limit, error) {
	return UpdateLimit(lc.c, limit)
}

func (lc *LimitClient) Delete(limit Limit) error {
	return DeleteLimit(lc.c, limit)
}

func (lc *LimitClient) DeleteById(limitId int64) error {
	return DeleteLimitById(lc.c, limitId)
}
