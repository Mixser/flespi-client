package flespi_subaccount

import "github.com/mixser/flespi-client/internal/flespiapi"

// SubaccountClient provides receiver-based methods for managing Flespi subaccounts.
// Access it via Client.Subaccounts after creating a flespi.Client.
type SubaccountClient struct {
	c flespiapi.Doer
}

// NewSubaccountClient creates a SubaccountClient wrapping the given flespiapi.Doer.
func NewSubaccountClient(c flespiapi.Doer) *SubaccountClient {
	return &SubaccountClient{c: c}
}

func (sc *SubaccountClient) Create(name string, options ...CreateSubaccountOption) (*Subaccount, error) {
	return NewSubaccount(sc.c, name, options...)
}

func (sc *SubaccountClient) List() ([]Subaccount, error) {
	return ListSubaccounts(sc.c)
}

func (sc *SubaccountClient) Get(subaccountId int64) (*Subaccount, error) {
	return GetSubaccount(sc.c, subaccountId)
}

func (sc *SubaccountClient) Update(subaccount Subaccount) (*Subaccount, error) {
	return UpdateSubaccount(sc.c, subaccount)
}

func (sc *SubaccountClient) Delete(subaccount Subaccount) error {
	return DeleteSubaccount(sc.c, subaccount)
}

func (sc *SubaccountClient) DeleteById(subaccountId int64) error {
	return DeleteSubaccountById(sc.c, subaccountId)
}
