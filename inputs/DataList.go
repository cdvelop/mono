package inputs

// eg: options=1:Admin,2:Editor,3:Visitante
func DataList(params ...any) datalist {
	new := datalist{
		attributes: attributes{
			htmlName: "datalist",
		},
	}
	new.Set(params)

	return new
}

type datalist struct {
	attributes
	dataSource
}

func (d datalist) ValidateInput(value string) error {
	return d.checkOptionKeys(value)
}

func (d datalist) BuildHtmlInput(id string) string {
	var req string
	if !d.allowSkipCompleted {
		req = ` required`
	}

	tag := `<input list="` + d.Name + `" name="` + d.Name + `" id="` + id + `"` + req + `>`
	tag += `<datalist id="` + id + `">`
	tag += d.GetAllTagOption()
	tag += `</datalist>`

	return tag
}

// retorna string html option de un select
func (d datalist) GetAllTagOption() (opt string) {

	for _, o := range d.options {
		for k, v := range o {
			opt += d.LabelOptSelect(k, v)
		}
	}

	return
}

// etiqueta html option de un datalist
func (d datalist) LabelOptSelect(key, value string) (opt string) {
	opt = `<option data-id="` + key + `" value="` + value + `">`
	return
}
