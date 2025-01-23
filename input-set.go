package mono

import (
	"fmt"
	"reflect"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type inputAdapter interface {
	Render(tabIndex *int) string
	Validate(value string) error
}

func (h *input) Set(params ...any) {
	if h.customName == "" {
		h.customName = h.htmlName
	}

	// h.searchDataSourceImplementation(params...)

	options := h.separateOptions(params...)

	for _, option := range options {
		switch option {
		case "hidden":
			h.htmlName = option
		case "!required":
			h.allowSkipCompleted = true
		case `typing="hide"`:
			h.htmlName = "password"
		case "multiple":
			h.Multiple = option
		case "letters":
			h.Letters = true
		case "numbers":
			h.Numbers = true
		}

		switch {

		case strings.Contains(option, "chars="):
			h.Characters = []rune(extractValue(option, "chars"))

		case strings.Contains(option, "data="):
			extractData(extractValue(option, "data"), &h.DataSet)

		case strings.Contains(option, "options="):
			extractData(extractValue(option, "options"), &h.options)

		case strings.Contains(option, "class="):
			newClass := className(extractValue(option, "class"))
			exists := false
			for _, class := range h.Class {
				if class == newClass {
					exists = true
					break
				}
			}
			if !exists {
				h.Class = append(h.Class, newClass)
			}
		case strings.Contains(option, "entity="):
			h.entity = extractValue(option, "entity")

		case strings.Contains(option, "name="):
			h.Name = extractValue(option, "name")

		case strings.Contains(option, "legend="):
			h.legend = extractValue(option, "legend")

		case strings.Contains(option, "min="):
			h.Min = extractValue(option, "min")

		case strings.Contains(option, "max="):
			h.Max = extractValue(option, "max")

		case strings.Contains(option, "maxlength="):
			h.Maxlength = extractValue(option, "maxlength")

		case strings.Contains(option, "placeholder="):
			h.PlaceHolder = extractValue(option, "placeholder")

		case strings.Contains(option, "title="):
			h.Title = extractValue(option, "title")

		case strings.Contains(option, "autocomplete="):
			h.Autocomplete = extractValue(option, "autocomplete")

		case strings.Contains(option, "rows="):
			h.Rows = extractValue(option, "rows")

		case strings.Contains(option, "cols="):
			h.Cols = extractValue(option, "cols")

		case strings.Contains(option, "step="):
			h.Step = extractValue(option, "step")

		case strings.Contains(option, "oninput="):
			h.Oninput = extractValue(option, "oninput")

		case strings.Contains(option, "onkeyup="):
			h.Onkeyup = extractValue(option, "onkeyup")

		case strings.Contains(option, "onchange="):
			h.Onchange = extractValue(option, "onchange")

		case strings.Contains(option, "value="):
			h.Value = extractValue(option, "value")

		case strings.Contains(option, "accept="):
			h.Accept = extractValue(option, "accept")

		}
	}

	if h.Name == "" {
		h.Name = h.customName
	}

	if h.htmlName != "hidden" {
		h.setDynamicTitle()
	} else {
		h.Title = ""
		h.PlaceHolder = ""
	}

	if len(h.options) == 0 {
		h.options = []map[string]string{
			{"": ""},
		}
	}

}

// Method to dynamically generate the title
func (h *input) setDynamicTitle() {
	if h.Title != "" {
		return
	}

	var parts []string
	parts = append(parts, R.T(D.Allowed))

	if h.Letters {
		parts = append(parts, R.T(D.Letters))
	}

	if h.Numbers {
		parts = append(parts, R.T(D.Numbers))
	}

	if len(h.Characters) > 0 {
		var chars []string
		for _, char := range h.Characters {
			if char == ' ' {
				chars = append(chars, "␣")
			} else {
				chars = append(chars, string(char))
			}
		}
		parts = append(parts, R.T(D.Chars, chars))
	}

	if h.Minimum != 0 {
		parts = append(parts, R.T("min", h.Minimum))
	}

	if h.Maximum != 0 {
		parts = append(parts, R.T("max", h.Maximum))
	}

	h.Title = strings.Join(parts, " ")
}

func (h *input) separateOptions(params ...any) (options []string) {
	for _, param := range params {
		h.processParam(param, &options)
	}
	return
}

func (h *input) processParam(param any, options *[]string) {
	switch v := param.(type) {
	case string:
		*options = append(*options, splitAndTrimOptions(v)...)
	case []string:
		for _, s := range v {
			*options = append(*options, splitAndTrimOptions(s)...)
		}

	case []any:
		for _, item := range v {
			h.processParam(item, options)
		}

	default:
		h.handleUnknownType(param)
	}
}

func (h *input) handleUnknownType(param any) {
	if t, ok := param.(reflect.Type); ok {
		// fmt.Println("*Tipo de reflect.Type:", t.Name())

		// Param is of type reflect.Type
		if sd, ok := reflect.New(t).Interface().(sourceData); ok {
			// fmt.Printf("El tipo %s implementa la interfaz sourceData\n", t.Name())
			h.data = sd
		}

		return
	}
	// Handle other unknown types or add default behavior

}
func splitAndTrimOptions(opt string) []string {
	var result []string
	if strings.Contains(opt, ";") {
		opts := strings.Split(opt, ";")
		for _, o := range opts {
			o = strings.TrimSpace(o)
			if o != "" {
				result = append(result, o)
			}
		}
	} else {
		opt = strings.TrimSpace(opt)
		if opt != "" {
			result = append(result, opt)
		}
	}
	return result
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

	inputType = G.String.SnakeCase(inputType)

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

		fmt.Println("f.Legend is empty: ", f.Name)

		name := R.T(f.Name)
		// f.Legend = cases.Title(language.Und).String(strings.ToLower(name))

		c := cases.Title(language.English)
		f.Legend = c.String(name)

	}

}
