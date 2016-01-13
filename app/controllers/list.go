package controllers

import (
	"github.com/revel/revel"
	"github.com/kwn/golang-lists/app/models"
	"github.com/kwn/golang-lists/app/errors"
	"github.com/kwn/golang-lists/app/forms"
	"fmt"
)

type List struct {
	*revel.Controller
}

func (c List) List() revel.Result {
	lists := []models.List{
		models.List{5, "Weekly shopping list", "weekly", "66", "pink", []models.ListItem{}, []models.ListProduct{}},
		models.List{6, "Other list", "daily", "66", "pink", []models.ListItem{}, []models.ListProduct{}},
	}

	return c.RenderJson(lists)
}

func (c List) Get(id int) revel.Result {
	if id != 5 {
		return c.RenderJson(errors.NewHttpErrorWithDetails(404, "List not found"))
	}

	list := models.List{5, "Weekly shopping list", "weekly", "66", "pink", []models.ListItem{}, []models.ListProduct{}}

	return c.RenderJson(list)
}

func (c List) Add() revel.Result {
	listForm := forms.List()
	form := listForm(c.Request.Request)

	if !form.IsValid() {
		return c.RenderJson(errors.NewHttpErrorWithDetails(400, fmt.Sprintf("%v", form.Errors())))
	}

	list := models.List{}
	form.MapTo(&list)

	return c.RenderJson(list)
}

func (c List) Delete() revel.Result {
	c.Response.Status = 204

	return c.RenderJson("")
}
