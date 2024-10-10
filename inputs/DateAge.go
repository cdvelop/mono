package inputs

// formato fecha: DD-MM-YYYY
// options: `title="xxx"`
func DateAge(params ...any) dateAge {
	new := dateAge{
		input: input{
			attributes: attributes{
				htmlName:   "date",
				customName: "DateAge",
				Title:      `title="formato fecha: DD-MM-YYYY"`,
				// Pattern: `[0-9]{4}-(0[1-9]|1[012])-(0[1-9]|1[0-9]|2[0-9]|3[01])`,
				// Onkeyup:  `onkeyup="DateAgeChange(this)"`,
				Onchange: `onchange="DateAgeChange(this)"`,
			},
		},
		day: Date(),
	}
	new.Set(params)

	return new
}

type dateAge struct {
	input
	day date
}

func (d dateAge) Render(id string) string {

	tag := `<label class="age-number"><input data-name="age-number" type="number" min="0" max="150" oninput="AgeInputChange(this)" title="Campo Informativo"></label>`

	tag += `<label class="age-date">`

	tag += d.input.Render(id)

	tag += `</label>`

	return tag
}

func (d dateAge) Validate(value string) error { //en realidad es YYYY-MM-DD
	return d.day.CheckDateExists(value)
}

func (d dateAge) GoodTestData() (out []string) {
	return d.day.GoodTestData()
}

func (d dateAge) WrongTestData() (out []string) {
	return d.day.WrongTestData()
}
