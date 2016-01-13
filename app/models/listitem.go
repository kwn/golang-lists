package models

type ListItem struct {
	ID     int    `json:"id"`
	Term   string `json:"term"`
	ListID List `json:"list" sql:"index"`
}
