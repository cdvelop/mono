package inputs

import (
	"fmt"
)

// eg: options=1:Admin,2:Editor,3:Visitante
func Select(params ...any) selec {

	new := selec{
		attributes: attributes{
			htmlName: "select",
		},
	}
	new.Set(params)
	return new
}

type selec struct {
	attributes
	dataSource
}

func (s selec) ValidateInput(value string) error {
	return s.checkOptionKeys(value)
}

func (s selec) BuildHtmlInput(id string) string {
	var req string
	if !s.allowSkipCompleted {
		req = ` required`
	}

	return fmt.Sprintf(`<select name="%v"%v><option selected></option>%v</select>`, s.Name, req, s.GetAllTagOption())
}

// retorna string html option de un selecTag
func (s selec) GetAllTagOption() (opt string) {

	for _, o := range s.options {
		for k, v := range o {
			opt += s.LabelOptSelect(k, v)
		}
	}

	return
}

// etiqueta html option de un selecTag [value=id name= texto a mostrar]
func (s selec) LabelOptSelect(key, value string) (opt string) {
	opt = `<option name="` + key + `" value="` + key + `">` + value + `</option>`
	return
}
