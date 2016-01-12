package controllers

import (
	"github.com/revel/revel"
	"github.com/kwn/golang-lists/app/models"
	"github.com/kwn/golang-lists/app/errors"
)

type List struct {
	*revel.Controller
}

func (c List) GetLists() revel.Result {
	lists := []models.JsList{
		models.JsList{5, "Weekly shopping list", "weekly", "66", "pink", []models.JsListItem{}, []models.JsListProduct{}},
		models.JsList{6, "Other list", "daily", "66", "pink", []models.JsListItem{}, []models.JsListProduct{}},
	}

	return c.RenderJson(lists)
}

func (c List) GetList(id int) revel.Result {
	if id != 5 {
		return c.RenderJson(errors.NewHttpErrorWithDetails(404, "List not found"))
	}

	list := models.JsList{5, "Weekly shopping list", "weekly", "66", "pink", []models.JsListItem{}, []models.JsListProduct{}}

	return c.RenderJson(list)
}

func (c List) PostList() revel.Result {

}
