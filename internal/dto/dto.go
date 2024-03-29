package dto

type UsdBrlDTO struct {
	Code       string `json:"code"`
	CodeIn     string `json:"code_in"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `jnson:"var_bid"`
	PctChange  string `json:"pct_change"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"tinmestamp"`
	CreateDate string `json:"create_date"`
}
