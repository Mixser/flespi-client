package flespi_subaccount

type Subaccount struct {
	Id      int64  `json:"id,omitempty"`
	Name    string `json:"name"`
	LimitId int64  `json:"limit_id"`

	Metadata map[string]string `json:"metadata,omitempty"`

	// AccountId is the parent subaccount that owns this subaccount (returned as "cid" in API responses).
	// On creation it is passed via the x-flespi-cid header, not the request body.
	AccountId int64 `json:"cid,omitempty"`
}

type CreateSubaccountOption func(*Subaccount)

func WithLimit(limitId int64) CreateSubaccountOption {
	return func(account *Subaccount) {
		account.LimitId = limitId
	}
}

func WithAccountId(accountId int64) CreateSubaccountOption {
	return func(account *Subaccount) {
		account.AccountId = accountId
	}
}

type subaccountsResponse struct {
	Subaccounts []Subaccount `json:"result"`
}
