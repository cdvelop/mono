package inputs

// eg: options=1:Admin,2:Editor,3:Visitante
func List(params ...any) list {
	new := list{
		input: input{
			attributes: attributes{
				htmlName: "ol",
			},
		},
	}
	new.Set(params)

	return new
}

type list struct {
	input
}

func (d list) Validate(value string) error {
	return nil
}

func (d list) Render(id string) string {

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
