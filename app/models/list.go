package models

type JsList struct {
	ID       int             `json:"id"`
	Name     string          `json:"name"`
	Type     string          `json:"type"`
	User     string          `json:"user"`
	Color    string          `json:"color"`
	Items    []JsListItem    `json:"items"`
	Products []JsListProduct `json:"products"`
}
