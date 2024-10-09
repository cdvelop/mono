package inputs

import (
	"errors"
	"reflect"
	"strings"
	"unicode"
)

type attributes struct {
	allowSkipCompleted bool //permite que el campo no sea completado

	Type     string // ej : text, password, email
	htmlName string //eg input,select,textarea
	Id       string // id="123"
	Name     string //eg address,phone

	PlaceHolder string
	Title       string //info

	Min string //valor mínimo
	Max string //valor máximo

	Maxlength string //ej: maxlength="12"

	Autocomplete string

	Rows string //filas ej 4,5,6
	Cols string //columnas ej 50,80

	Step     string
	Oninput  string // ej: "miRealtimeFunction()" = oninput="miRealtimeFunction()"
	Onkeyup  string // ej: "miNormalFuncion()" = onkeyup="miNormalFuncion()"
	Onchange string // ej: "miNormalFuncion()" = onchange="myFunction()"

	// https://developer.mozilla.org/en-US/docs/Web/HTML/attributes/accept
	// https://developer.mozilla.org/es/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types
	// accept="image/*"
	Accept   string
	Multiple string // multiple

	Value string

	customName string //eg onlyText,onlyNumber...

	Class []string // clase css ej: class="age"

	DataSet []map[string]string // dataset ej: data-id="123" = map[string]string{"id": "123"}

	options []map[string]string // ej: [{"m": "male"}, { "f": "female"}]

	list []string // ej: ["a","b","c"]
}

// extractValue removes a prefix. Example:
// `min="5"` and delete="min", it returns "5"
func extractValue(option, delete string) string {
	out := strings.Replace(option, delete+`="`, "", 1)
	if strings.Contains(out, delete) {
		out = strings.Replace(option, delete+`=`, "", 1)
	}
	// fmt.Println("option:", option, "delete:", delete, "out:", out)
	out = strings.TrimSuffix(out, `"`)
	return out
}

func (a *attributes) Set(params ...any) {
	if a.customName == "" {
		a.customName = a.htmlName
	}

	ptd, options := a.separateOptions(params...)

	for _, option := range options {
		switch option {

		case "hidden":
			a.htmlName = option

		case "!required":
			a.allowSkipCompleted = true

		case `typing="hide"`:
			a.htmlName = "password"

		case "multiple":
			a.Multiple = option

		case "letters":
			if ptd != nil {
				ptd.Letters = true
			}
		case "numbers":
			if ptd != nil {
				ptd.Numbers = true
			}
		}

		switch {

		case strings.Contains(option, "chars="):
			ptd.Characters = []rune(extractValue(option, "chars"))

		case strings.Contains(option, "data="):
			extractData(extractValue(option, "data"), &a.DataSet)

		case strings.Contains(option, "options="):
			extractData(extractValue(option, "options"), &a.options)

		case strings.Contains(option, "class="):
			a.Class = append(a.Class, extractValue(option, "class"))

		case strings.Contains(option, "name="):
			a.Name = extractValue(option, "name")

		case strings.Contains(option, "min="):
			a.Min = extractValue(option, "min")

		case strings.Contains(option, "max="):
			a.Max = extractValue(option, "max")

		case strings.Contains(option, "maxlength="):
			a.Maxlength = extractValue(option, "maxlength")

		case strings.Contains(option, "placeholder="):
			a.PlaceHolder = extractValue(option, "placeholder")

		case strings.Contains(option, "title="):
			a.Title = extractValue(option, "title")

		case strings.Contains(option, "autocomplete="):
			a.Autocomplete = extractValue(option, "autocomplete")

		case strings.Contains(option, "rows="):
			a.Rows = extractValue(option, "rows")

		case strings.Contains(option, "cols="):
			a.Cols = extractValue(option, "cols")

		case strings.Contains(option, "step="):
			a.Step = extractValue(option, "step")

		case strings.Contains(option, "oninput="):
			a.Oninput = extractValue(option, "oninput")

		case strings.Contains(option, "onkeyup="):
			a.Onkeyup = extractValue(option, "onkeyup")

		case strings.Contains(option, "onchange="):
			a.Onchange = extractValue(option, "onchange")

		case strings.Contains(option, "value="):
			a.Value = extractValue(option, "value")

		case strings.Contains(option, "accept="):
			a.Accept = extractValue(option, "accept")

		}
	}

	if ptd != nil && a.Title == "" {
		ptd.setDynamicTitle(a)
	}

}

func (a attributes) BuildHtmlInput(id string) string {
	return a.buildHtml(id)
}

func (a attributes) buildHtml(id string) (result string) {
	var open = `<input`
	var close = `>`
	a.Type = a.htmlName

	if a.htmlName == "textarea" {
		open = `<textarea`
		close = `></textarea>`
		a.Type = ""
	}

	a.Id = id

	a.DataSet = append(a.DataSet, map[string]string{
		"name": a.customName,
	})

	result = open

	elem := reflect.ValueOf(a)
	elemType := elem.Type()

	for i := 0; i < elem.NumField(); i++ {
		rf := elem.Field(i)
		attributeName := elemType.Field(i).Name

		// skip if it starts with lowercase
		if !unicode.IsUpper([]rune(attributeName)[0]) {
			continue
		}

		switch rf.Kind() {
		case reflect.String:
			fieldValue := rf.String()

			if fieldValue != "" {
				htmlAttribute(&result, attributeName, fieldValue)
			}

		case reflect.Slice:
			if attributeName == "DataSet" {
				for _, dataAttr := range rf.Interface().([]map[string]string) {
					for key, value := range dataAttr {
						htmlAttribute(&result, "data-"+key, value)
					}
				}
			}
		}
	}

	if !a.allowSkipCompleted {
		result += ` required`
	}

	result += close

	return result
}

func htmlAttribute(out *string, key, value string) {
	key = strings.ToLower(key)
	*out += ` ` + key + `="` + value + `"`
}

// extract options eg: "options=m:male,f:female" to:
// []map[string]string{{"m":"male"},{"f":"female"}}
func extractData(dataIn string, out *[]map[string]string) {

	options := strings.Split(dataIn, ",")

	for _, option := range options {
		parts := strings.Split(option, ":")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			*out = append(*out, map[string]string{key: value})
		}
	}
}
func (a attributes) checkOptionKeys(value string) error {

	dataInArray := strings.Split(value, ",")

	for _, keyIn := range dataInArray {

		if keyIn == "" {
			return errors.New("selección requerida campo " + a.Name)
		}

		var exist bool
		// fmt.Println("a.optionKeys", a.optionKeys)
		for _, opt := range a.options {
			if _, exist = opt[keyIn]; exist {
				break
			}
		}

		if !exist {
			return errors.New("valor " + keyIn + " no permitido en " + a.htmlName + " " + a.Name)
		}

	}

	return nil

}

func (a attributes) GoodTestData() (out []string) {
	for _, opt := range a.options {
		for k := range opt {
			out = append(out, k)
		}
	}
	return
}

func (a attributes) WrongTestData() (out []string) {
	for _, wd := range wrong_data {
		found := false
		for _, opt := range a.options {
			if _, exists := opt[wd]; exists {
				found = true
				break
			}
		}
		if !found {
			out = append(out, wd)
		}
	}
	return
}

func (a *attributes) separateOptions(params ...any) (ptd *permitted, options []string) {
	for _, param := range params {
		ptd, options = a.processParam(param, ptd, options)
	}
	return
}

func (a *attributes) processParam(param any, ptd *permitted, options []string) (*permitted, []string) {
	switch v := param.(type) {
	case string:
		options = append(options, a.splitAndTrimOptions(v)...)
	case []string:
		for _, s := range v {
			options = append(options, a.splitAndTrimOptions(s)...)
		}
	case *permitted:
		ptd = v
	case []any:
		for _, item := range v {
			ptd, options = a.processParam(item, ptd, options)
		}
	}
	return ptd, options
}

func (a *attributes) splitAndTrimOptions(opt string) []string {
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
