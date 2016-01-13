package models

type ListProduct struct {
	ID       int    `json:"id"`
	Term     string `json:"term"`
	Sku      string `json:"sku"`
	Quantity int    `json:"quantity"`
	ListID   List `json:"list" sql:"index"`
}
