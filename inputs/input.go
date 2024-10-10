package inputs

import (
	"fmt"
	"strings"
)

type input struct {
	attributes
	permitted
	dataSource
}

func (h *input) Set(params ...any) {
	if h.customName == "" {
		h.customName = h.htmlName
	}

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
			h.Class = append(h.Class, extractValue(option, "class"))

		case strings.Contains(option, "name="):
			h.Name = extractValue(option, "name")

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

	h.setDynamicTitle()

}

// Método para generar dinámicamente el título
func (h *input) setDynamicTitle() {

	if h.Title != "" {
		return
	}

	var parts []string
	parts = append(parts, "permitido:")

	// Lógica de validación para letras
	if h.Letters {
		parts = append(parts, "letras")
	}

	// Lógica de validación para números
	if h.Numbers {
		parts = append(parts, "números")
	}

	// Lógica de validación para caracteres permitidos
	if len(h.Characters) > 0 {
		var chars []string
		for _, char := range h.Characters {
			if char == ' ' {
				chars = append(chars, "␣") // Reemplaza el espacio con el carácter visible '␣'
			} else {
				chars = append(chars, string(char))
			}
		}
		parts = append(parts, fmt.Sprintf("caracteres: %v ", strings.Join(chars, " ")))
	}

	if h.Minimum != 0 {
		parts = append(parts, fmt.Sprintf("min. %d", h.Minimum))
	}

	if h.Maximum != 0 {
		parts = append(parts, fmt.Sprintf("max. %d", h.Maximum))
	}

	// Generar el valor final para el atributo Title
	h.Title = strings.Join(parts, " ")
}

func (h *input) separateOptions(params ...any) (options []string) {
	for _, param := range params {
		options = h.processParam(param, options)
	}
	return
}

func (h *input) processParam(param any, options []string) []string {
	switch v := param.(type) {
	case string:
		options = append(options, splitAndTrimOptions(v)...)
	case []string:
		for _, s := range v {
			options = append(options, splitAndTrimOptions(s)...)
		}

	case []any:
		for _, item := range v {
			options = h.processParam(item, options)
		}
	}
	return options
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
