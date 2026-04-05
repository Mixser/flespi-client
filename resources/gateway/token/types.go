package flespi_token

type Token struct {
	Id  int64  `json:"id,omitempty"`
	Key string `json:"key,omitempty"`

	Info string `json:"info"`

	Enabled bool `json:"enabled"`

	Expire int64 `json:"expire"`
	TTL    int64 `json:"ttl"`

	SubAccountId int64             `json:"cid,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}

type CreateTokenOption func(*Token)

func WithStatus(enabled bool) CreateTokenOption {
	return func(token *Token) {
		token.Enabled = enabled
	}
}

func WithExpire(expire int64) CreateTokenOption {
	return func(token *Token) {
		token.Expire = expire
	}
}

func WithTTL(ttl int64) CreateTokenOption {
	return func(token *Token) {
		token.TTL = ttl
	}
}

func WithSubAccountId(subAccountId int64) CreateTokenOption {
	return func(token *Token) {
		token.SubAccountId = subAccountId
	}
}

func WithMetadata(metadata map[string]string) CreateTokenOption {
	return func(token *Token) {
		token.Metadata = metadata
	}
}

type tokensResponse struct {
	Tokens []Token `json:"result"`
}
