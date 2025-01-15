package inputs

import "strconv"

// eg: options=1:Admin,2:Editor,3:Visitante
func CheckBox(params ...any) *check {
	new := &check{
		input: input{
			attributes: attributes{
				htmlName: "checkbox",
			},
		},
	}
	new.Set(params)
	return new
}

type check struct {
	input
	onlyInternalContend bool
}

func (c check) Validate(value string) error {
	return c.checkOptionKeys(value)
}

func (c check) Render(tabIndex int) string {
	var tags string
	for i, opt := range c.options {
		id3 := strconv.Itoa(tabIndex) + "." + strconv.Itoa(i)
		var field_value string
		var text_field string

		for k, v := range opt {
			field_value = k
			text_field = v
		}

		tag_input := `<input type="checkbox" id="` + id3 + `" name="` + c.Name + `" value="` + field_value + `" onchange="CheckChange(this)"><span>` + text_field + `</span>`

		if !c.onlyInternalContend {
			tags += `<label data-id="` + field_value + `" for="` + id3 + `" class="block-label">` + tag_input + `</label>`
		} else {
			tags += tag_input
		}
	}

	return tags
}
