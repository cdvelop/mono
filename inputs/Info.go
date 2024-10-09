package inputs

//value="valor a mostrar"
func Info(params ...any) info {
	new := info{
		attributes: attributes{
			htmlName:   "text",
			customName: "info",
		},
	}
	new.Set(params)
	return new
}

// input de carácter informativo
type info struct {
	attributes
}

func (i info) BuildHtmlInput(id string) string {
	return i.Value
}

func (i info) ValidateInput(string) error {
	return nil
}
