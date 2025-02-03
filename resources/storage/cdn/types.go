package flespi_cdn


type CDN struct {
	Id      int64    `json:"id,omitempty"`
	Name    string `json:"name"`
	Blocked bool   `json:"blocked,omitempty"`
	Size    int64    `json:"size,omitempty"`
}


type cdnsResponse struct{
	CDNS []CDN `json:"result"`
}


type CreateCDNOption func(*CDN)
