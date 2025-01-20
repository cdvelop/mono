package inputs

import (
	"reflect"
	"strings"
)

type input struct {
	attributes
	permitted
	dataSource
	fieldset
}

type fieldset struct {
	cssClasses []className
}

type container struct {
	cssClasses []className
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
	parts = append(parts, Lang.T(D.Allowed))

	if h.Letters {
		parts = append(parts, Lang.T(D.Letters))
	}

	if h.Numbers {
		parts = append(parts, Lang.T(D.Numbers))
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
		parts = append(parts, Lang.T(D.Chars, chars))
	}

	if h.Minimum != 0 {
		parts = append(parts, Lang.T("min", h.Minimum))
	}

	if h.Maximum != 0 {
		parts = append(parts, Lang.T("max", h.Maximum))
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
