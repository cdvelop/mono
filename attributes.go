package mono

import (
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

type attributes struct {
	allowSkipCompleted bool //permite que el campo no sea completado
	hideTyping         bool //oculta el valor mientras se escribe

	Type     string // input type  eg : text, password, email
	htmlName string //eg input,select,textarea
	Name     string //eg address,phone
	Id       string
	entity   string //eg: user, product, order
	legend   string

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

	Class []className // clase css ej: class="age"

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

func (h input) Render(tabIndex *int) string {
	var tags, close, labelForID string

	for _, opt := range h.options {
		h.Id = h.htmlID() //generate new id
		if labelForID == "" {
			labelForID = h.Id
		}

		for value, text := range opt {

			switch h.htmlName {

			case "datalist":
				if tags == "" {
					h.htmlName = "list"
					tags += h.renderOneInput()
					tags += `<datalist id="` + labelForID + `">`
					tags += `<option value="` + text + `">`
					close = `</datalist>`
					h.htmlName = "datalist"
				}

				tags += `<option value="` + text + `">`
			case "list":
				if tags == "" {
					tags += `<ol>`
					tags += `<li><p>` + value + `: ` + text + `</p></li>`
					close = `</ol>`
				}

				tags += `<li><p>` + value + `: ` + text + `</p></li>`

			case "checkbox", "radio":
				h.Value = value
				tags += `<label for="` + h.Id + `">`
				tags += h.renderOneInput()
				tags += `<span>` + text + `</span>`
				tags += `</label>`
				h.allowSkipCompleted = true // avoid required appearing again
				*tabIndex++

			case "select":
				if tags == "" {
					tags += h.renderOneInput()
					tags += `<option ` + R.T(D.Select) + `></option>`
					close = `</select>`
				}
				tags = `<option value="` + value + `">` + text + `</option>`
			default:
				tags += h.renderOneInput()
			}
		}
	}

	tags += close

	return h.renderInputLayout(tags, labelForID, tabIndex)
}

func (h input) renderOneInput() (result string) {
	var open = `<input`
	var close = `>`
	h.Type = h.htmlName

	switch h.htmlName {
	case "textarea":
		open = `<textarea`
		close = `></textarea>`
		h.Type = ""
	case "select":
		open = `<select`
		h.Type = ""
	}

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

func (h input) renderInputLayout(inputHtml, labelForID string, tabindex *int) string {

	if h.htmlName == "hidden" {
		return inputHtml
	}

	return `<fieldset tabindex="` + strconv.Itoa(*tabindex) + `"` + h.getClassNames() + `">
	<legend><label for="` + labelForID + `">` + h.legend + `</label></legend>
	` + inputHtml + `
</fieldset>`
}

// getClassNames retorna los nombres de múltiples clases concatenados para HTML
func (h input) getClassNames() string {
	var names []string
	for _, className := range h.cssClasses {
		if className != "" {
			names = append(names, string(className))
		}
	}

	if len(names) == 0 {
		return ""
	}
	return ` class="` + strings.Join(names, " ") + `"`
}

var inputId int

func (h input) htmlID() string {

	if h.Id != "" {
		return h.Id
	}

	inputId++
	return strconv.Itoa(inputId)
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
