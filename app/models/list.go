package models

type List struct {
	ID       int           `json:"id"`
	Name     string        `json:"name" gforms:"name"`
	Type     string        `json:"type" gforms:"type"`
	User     string        `json:"user"`
	Color    string        `json:"color" gforms:"color"`
	Items    []ListItem    `json:"items"`
	Products []ListProduct `json:"products"`
}
