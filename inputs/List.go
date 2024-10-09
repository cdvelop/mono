package inputs

// eg: options=1:Admin,2:Editor,3:Visitante
func List(params ...any) list {
	new := list{
		attributes: attributes{
			htmlName: "ol",
		},
	}
	new.Set(params)

	return new
}

type list struct {
	attributes
	dataSource
}

func (d list) ValidateInput(value string) error {
	return nil
}

func (d list) BuildHtmlInput(id string) string {

	tag := `<ol>`
	tag += d.getAll()
	tag += `</ol>`

	return tag
}

func (h list) getAll() (opt string) {
	for _, o := range h.options {
		opt += `<li>`
		for k, v := range o {
			opt += `<p>` + k + `: ` + v + `</p>`
		}
		opt += `</li>`
	}
	return
}
