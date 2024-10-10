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

func (h input) Render(id string) (result string) {
	var open = `<input`
	var close = `>`
	h.Type = h.htmlName

	if h.htmlName == "textarea" {
		open = `<textarea`
		close = `></textarea>`
		h.Type = ""
	}

	h.Id = id

	h.DataSet = append(h.DataSet, map[string]string{
		"name": h.customName,
	})

	result = open

	elem := reflect.ValueOf(h.attributes)
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

	if !h.allowSkipCompleted {
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
