package models

type JsListItem struct {
	ID     int    `json:"id"`
	Term   string `json:"term"`
	ListID JsList `json:"list" sql:"index"`
}
