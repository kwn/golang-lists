package forms

import "github.com/bluele/gforms"

func List() gforms.Form {
	return gforms.DefineForm(gforms.NewFields(
		gforms.NewTextField(
			"type",
			gforms.Validators{
				gforms.Required(),
				gforms.MaxLengthValidator(255),
			},
		),
		gforms.NewTextField(
			"name",
			gforms.Validators{
				gforms.Required(),
				gforms.MaxLengthValidator(255),
			},
		),
		gforms.NewTextField(
			"color",
			gforms.Validators{
				gforms.Required(),
				gforms.MaxLengthValidator(255),
				gforms.RegexpValidator(`(white|black|light_pink|pink)`),
			},
		),
	))
}
