package models

type JsListProduct struct {
	ID       int    `json:"id"`
	Term     string `json:"term"`
	Sku      string `json:"sku"`
	Quantity int    `json:"quantity"`
	ListID   JsList `json:"list" sql:"index"`
}
