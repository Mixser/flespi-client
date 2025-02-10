package flespi_subaccount

type Subaccount struct {
	Id      int64  `json:"id,omitempty"`
	Name    string `json:"name"`
	LimitId int64  `json:"limit_id"`

	Metadata map[string]string `json:"metadata,omitempty"`
}

type CreateSubaccountOption func(*Subaccount)

func WithLimit(limitId int64) CreateSubaccountOption {
	return func(account *Subaccount) {
		account.LimitId = limitId
	}
}

type subaccountsResponse struct {
	Subaccounts []Subaccount `json:"result"`
}
