package mono

import (
	"reflect"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type inputAdapter interface {
	Render(tabIndex *int) string
	Validate(value string) error
}

func (f *field) setInput(structureFrom reflect.Type, rf *reflect.StructField) {

	var params []any

	if structureFrom != nil {
		params = append(params, structureFrom)
	}

	inputData := rf.Tag.Get("Input")
	var inputType string
	if inputData != "" {

		// Get base params from input tag
		params = append(params, getParams(inputData))

		// Get input type from input tag
		inputType = strings.Split(inputData, "(")[0]
	}

	if inputType == "" {
		// set input type by field name
		inputType = f.Name
		// fmt.Println("SET inputType ", inputType)
	}

	inputType = snakeCase(inputType)

	// Add common params
	params = append(params,
		"name="+f.Name,
		"entity="+f.Parent.Name,
	)

	if f.Legend != "" {
		params = append(params, "legend="+f.Legend)
	}

	// fmt.Println("setInput params:", params)

	// Handle specific input types
	switch inputType {
	case "checkbox":
		f.Input = IN.CheckBox(params...)
	case "datalist":
		f.Input = IN.DataList(params...)
	case "date", "birth_date":
		f.Input = IN.Date(params...)
	case "date_age":
		f.Input = IN.DateAge(params...)
	case "day_word":
		f.Input = IN.DayWord(params...)
	case "file_path":
		f.Input = IN.FilePath(params...)
	case "hour":
		f.Input = IN.Hour(params...)
	case "id":
		f.Input = IN.ID(params...)
	case "info":
		f.Input = IN.Info(params...)
	case "ip":
		f.Input = IN.Ip(params...)
	case "list":
		f.Input = IN.List(params...)
	case "mail":
		f.Input = IN.Mail(params...)
	case "month_day":
		f.Input = IN.MonthDay(params...)
	case "number":
		f.Input = IN.Number(params...)
	case "password":
		f.Input = IN.Password(params...)
	case "phone":
		f.Input = IN.Phone(params...)
	case "radio":
		f.Input = IN.Radio(params...)
	case "gender":
		f.Input = IN.RadioGender(params...)
	case "rut":
		f.Input = IN.Rut(params...)
	case "select":
		if structureFrom != nil {
			params = append(params, "structure="+structureFrom.String())
		}
		f.Input = IN.Select(params...)
	case "text", "name":
		f.Input = IN.Text(params...)
	case "text_area":
		f.Input = IN.TextArea(params...)
	case "text_number":
		f.Input = IN.TextNumber(params...)
	case "text_number_code":
		f.Input = IN.TextNumberCode(params...)
	case "text_only":
		f.Input = IN.TextOnly(params...)
	case "text_search":
		f.Input = IN.TextSearch(params...)
	}
}

func getParams(inputTag string) []string {
	start := strings.Index(inputTag, "(")
	end := strings.Index(inputTag, ")")
	if start == -1 || end == -1 {
		return nil
	}
	paramsStr := inputTag[start+1 : end]
	return strings.Split(paramsStr, ",")
}

func (f *field) setLegend(rf *reflect.StructField) {

	f.Legend = rf.Tag.Get("Legend")
	if f.Legend == "" {

		name := R.T(f.Name)
		// f.Legend = cases.Title(language.Und).String(strings.ToLower(name))

		c := cases.Title(language.English)
		f.Legend = c.String(name)

	}

}
