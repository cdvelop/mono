package inputs

// formato dia DD como palabra ej. Lunes 24 Diciembre
// options: title="xxx"
func DayWord(params ...any) dayWord {
	new := dayWord{
		month: MonthDay(),
		attributes: attributes{
			htmlName:   "text",
			customName: "DayWord",
			DataSet:    []map[string]string{{"spanish": ""}},
			// Pattern: `^[0-9]{2,2}$`,
		},
	}
	new.Set(params)

	return new
}

type dayWord struct {
	month monthDay
	attributes
}

func (d dayWord) InputName(customName, htmlName *string) {
	if customName != nil {
		*customName = "DayWord"
	}
	if htmlName != nil {
		*htmlName = "text"
	}
}

func (d dayWord) BuildHtmlInput(id string) string {
	tag := `<label class="date-spanish">`
	tag += d.buildHtml(id)
	tag += `</label>`
	return tag
}

func (d dayWord) ValidateInput(value string) error {
	return d.month.ValidateInput(value)
}

func (d dayWord) GoodTestData() (out []string) {
	return d.month.GoodTestData()
}

func (d dayWord) WrongTestData() (out []string) {
	return d.month.WrongTestData()
}
