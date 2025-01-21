package mono

import (
	"reflect"
	"strings"

	"github.com/cdvelop/mono/inputs"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type input interface {
	Render(tabIndex int) string
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
		f.Input = inputs.CheckBox(params...)
	case "datalist":
		f.Input = inputs.DataList(params...)
	case "date", "birth_date":
		f.Input = inputs.Date(params...)
	case "date_age":
		f.Input = inputs.DateAge(params...)
	case "day_word":
		f.Input = inputs.DayWord(params...)
	case "file_path":
		f.Input = inputs.FilePath(params...)
	case "hour":
		f.Input = inputs.Hour(params...)
	case "id":
		f.Input = inputs.ID(params...)
	case "info":
		f.Input = inputs.Info(params...)
	case "ip":
		f.Input = inputs.Ip(params...)
	case "list":
		f.Input = inputs.List(params...)
	case "mail":
		f.Input = inputs.Mail(params...)
	case "month_day":
		f.Input = inputs.MonthDay(params...)
	case "number":
		f.Input = inputs.Number(params...)
	case "password":
		f.Input = inputs.Password(params...)
	case "phone":
		f.Input = inputs.Phone(params...)
	case "radio":
		f.Input = inputs.Radio(params...)
	case "gender":
		f.Input = inputs.RadioGender(params...)
	case "rut":
		f.Input = inputs.Rut(params...)
	case "select":
		if structureFrom != nil {
			params = append(params, "structure="+structureFrom.String())
		}
		f.Input = inputs.Select(params...)
	case "text", "name":
		f.Input = inputs.Text(params...)
	case "text_area":
		f.Input = inputs.TextArea(params...)
	case "text_number":
		f.Input = inputs.TextNumber(params...)
	case "text_number_code":
		f.Input = inputs.TextNumberCode(params...)
	case "text_only":
		f.Input = inputs.TextOnly(params...)
	case "text_search":
		f.Input = inputs.TextSearch(params...)
		f.Input = inputs.Text(params...)
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

		name := Lang.T(f.Name)
		// f.Legend = cases.Title(language.Und).String(strings.ToLower(name))

		c := cases.Title(language.English)
		f.Legend = c.String(name)

	}

}
